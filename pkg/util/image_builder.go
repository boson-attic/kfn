package util

import (
	"context"
	"fmt"
	"github.com/containers/buildah"
	"github.com/containers/buildah/pkg/unshare"
	"github.com/containers/image/transports"
	"github.com/containers/image/types"
	"github.com/containers/storage"
	"github.com/opencontainers/go-digest"
	log "github.com/sirupsen/logrus"
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
		Isolation:        config.BuildahIsolation,
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
		log.Infof("Copying into container image %s to %s", add.From, add.To)
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
		Isolation: config.BuildahIsolation,
	}
	for _, cmd := range commands {
		log.Infof("Running command %s in directory %s", cmd.Command, cmd.Wd)

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

func CommitImage(builder *buildah.Builder, ctx *types.SystemContext, imageName string, imageTag string) (image.FunctionImage, error) {
	img := image.FunctionImage{
		ImageName: imageName,
		Tag:       imageTag,
	}

	log.Debugf("Using FunctionImage %+v", img)

	imageRef, err := img.ParseSpecDest()

	log.Infof("Commiting image to %s", transports.ImageName(imageRef))

	if err != nil {
		return image.FunctionImage{}, err
	}

	_, _, _, err = builder.Commit(context.TODO(), imageRef, buildah.CommitOptions{
		PreferredManifestType: buildah.Dockerv2ImageManifest,
		SystemContext:         ctx,
	})

	return img, err
}
