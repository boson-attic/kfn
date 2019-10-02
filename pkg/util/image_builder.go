package util

import (
	"context"
	"fmt"
	"github.com/containers/buildah"
	"github.com/containers/buildah/pkg/unshare"
	"github.com/containers/image/types"
	"github.com/containers/storage"
	"github.com/opencontainers/go-digest"
	"github.com/slinkydeveloper/kfn/pkg/config"
	"github.com/slinkydeveloper/kfn/pkg/image"
	"strings"
)

var digester = digest.Canonical.Digester()

func InitializeBuilder(ctx context.Context, systemContext *types.SystemContext, fromImage string) (*buildah.Builder, error) {
	buildStoreOptions, err := storage.DefaultStoreOptions(unshare.IsRootless(), unshare.GetRootlessUID())

	if err != nil {
		return nil, err
	}

	buildStore, err := storage.GetStore(buildStoreOptions)

	if err != nil {
		return nil, err
	}

	buildOpts := &buildah.CommonBuildOptions{}

	opts := buildah.BuilderOptions{
		FromImage:        fromImage,
		Isolation:        config.GetBuildahIsolation(),
		CommonBuildOpts:  buildOpts,
		ConfigureNetwork: buildah.NetworkDefault,
		SystemContext:    systemContext,
		Format:           buildah.Dockerv2ImageManifest,
	}

	return buildah.NewBuilder(ctx, buildStore, opts)
}

type BuildAdd struct {
	From string
	To   string
}

func Add(builder *buildah.Builder, adds ...BuildAdd) error {
	for _, add := range adds {
		err := builder.Add(add.To, false, buildah.AddAndCopyOptions{Hasher: digester.Hash()}, add.From)
		if err != nil {
			return fmt.Errorf("error while adding: %v", err)
		}
	}
	return nil
}

type BuildCommand struct {
	Command string
	Wd      string
}

func RunCommands(builder *buildah.Builder, commands ...BuildCommand) error {
	logger := config.GetLoggerWriter()
	runOptions := buildah.RunOptions{
		Stdout:    logger,
		Stderr:    logger,
		Isolation: config.GetBuildahIsolation(),
	}
	for _, cmd := range commands {
		command := strings.Split(cmd.Command, " ")

		if cmd.Wd != "" {
			runOptions.WorkingDir = cmd.Wd
		}

		if err := builder.Run(command, runOptions); err != nil {
			return fmt.Errorf("error while runnning command: %v", err)
		}
	}
	return nil
}

func CommitImage(builder *buildah.Builder, imageName string, imageTag string) (image.FunctionImage, error) {
	img := image.FunctionImage{
		ImageName: imageName,
		Tag:       imageTag,
	}

	imageRef, err := img.ParseSpecDest()

	if err != nil {
		return image.FunctionImage{}, err
	}

	_, _, _, err = builder.Commit(context.TODO(), imageRef, buildah.CommitOptions{
		PreferredManifestType: buildah.Dockerv2ImageManifest,
	})

	return img, err
}
