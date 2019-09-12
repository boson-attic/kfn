package image

import (
	"context"
	"fmt"
	"github.com/containers/buildah"
	"github.com/containers/buildah/pkg/unshare"
	"github.com/containers/image/transports"
	"github.com/containers/image/transports/alltransports"
	"github.com/containers/image/types"
	"github.com/containers/storage"
	"strings"
)

func (image FunctionImage) PushImage(systemContext *types.SystemContext, imageId string) error {
	dest, err := image.parseSpecDest()

	if err != nil {
		return err
	}

	buildStoreOptions, err := storage.DefaultStoreOptions(unshare.IsRootless(), unshare.GetRootlessUID())

	if err != nil {
		return err
	}

	buildStore, err := storage.GetStore(buildStoreOptions)

	if err != nil {
		return err
	}

	options := buildah.PushOptions{
		Store:         buildStore,
		SystemContext: systemContext,
	}

	st, st2, err := buildah.Push(context.TODO(), imageId, dest, options)

	fmt.Printf("%+v\n%+v\n", st, st2)

	return err
}

func (image FunctionImage) parseSpecDest() (types.ImageReference, error) {
	destSpec := fmt.Sprintf("%s/%s:%s", image.ImageRegistry, image.ImageName, image.Tag)
	dest, err := alltransports.ParseImageName(destSpec)
	// add the docker:// transport to see if they neglected it.
	if err != nil {
		destTransport := strings.Split(destSpec, ":")[0]
		if t := transports.Get(destTransport); t != nil {
			return dest, nil
		}

		if strings.Contains(destSpec, "://") {
			return dest, nil
		}

		destSpec = "docker://" + destSpec
		dest2, err2 := alltransports.ParseImageName(destSpec)
		if err2 != nil {
			return dest, nil
		}
		dest = dest2
	}
	return dest, nil
}
