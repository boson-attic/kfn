package image

import (
	"context"
	"github.com/containers/buildah"
	"github.com/containers/buildah/pkg/unshare"
	"github.com/containers/image/transports/alltransports"
	"github.com/containers/image/types"
	"github.com/containers/storage"
	"github.com/opencontainers/go-digest"
	"os"
	"path"
	"strings"
)

type BuildCommand struct {
	Command string
	Wd      string
}

func (image FunctionImage) BuildImage(ctx *types.SystemContext, targetDir string) (string, error) {
	buildStoreOptions, err := storage.DefaultStoreOptions(unshare.IsRootless(), unshare.GetRootlessUID())

	if err != nil {
		return "", err
	}

	buildStore, err := storage.GetStore(buildStoreOptions)

	if err != nil {
		return "", err
	}

	buildOpts := &buildah.CommonBuildOptions{}

	var isolation buildah.Isolation

	envIsolation := os.Getenv("BUILDAH_ISOLATION")

	switch envIsolation {
	case "chroot":
		isolation = buildah.IsolationChroot
	default:
		isolation = buildah.IsolationOCIRootless
	}

	opts := buildah.BuilderOptions{
		FromImage:        "node:12-alpine",
		Registry:         image.ImageRegistry,
		Isolation:        isolation,
		CommonBuildOpts:  buildOpts,
		ConfigureNetwork: buildah.NetworkDefault,
		SystemContext:    ctx,
		Format:           buildah.Dockerv2ImageManifest,
	}

	builder, err := buildah.NewBuilder(context.TODO(), buildStore, opts)

	if err != nil {
		return "", err
	}

	builder.SetPort("8080")

	digester := digest.Canonical.Digester()

	err = builder.Add("/home/node/src", false, buildah.AddAndCopyOptions{Hasher: digester.Hash()}, path.Join(targetDir, "src"))
	if err != nil {
		return "", err
	}

	err = builder.Add("/home/node/usr", false, buildah.AddAndCopyOptions{Hasher: digester.Hash()}, path.Join(targetDir, "usr"))
	if err != nil {
		return "", err
	}

	err = runCommands(builder, []BuildCommand{{
		"mkdir -p /home/node/usr/.npm", "",
	}, {
		"chmod -R a+g+x /home/node/usr", "",
	}, {
		"chmod -R a+g+x /home/node/src", "",
	}, {
		"npm install", "/home/node/usr",
	}, {
		"npm install", "/home/node/src",
	}})
	if err != nil {
		return "", err
	}

	builder.SetEnv("HOME", "/home/node/usr")
	builder.SetUser("1000")
	builder.SetWorkDir("/home/node/src")

	builder.SetCmd([]string{"node", "/home/node/src/index.js"})

	imageRef, err := alltransports.ParseImageName(image.ImageName)

	imageId, _, _, err := builder.Commit(context.TODO(), imageRef, buildah.CommitOptions{
		PreferredManifestType: buildah.Dockerv2ImageManifest,
	})

	return imageId, err
}

func runCommands(builder *buildah.Builder, commands []BuildCommand) error {
	for _, cmd := range commands {
		command := strings.Split(cmd.Command, " ")
		runOptions := buildah.RunOptions{}
		if cmd.Wd != "" {
			runOptions.WorkingDir = cmd.Wd
		}

		if err := builder.Run(command, runOptions); err != nil {
			return err
		}
	}
	return nil
}
