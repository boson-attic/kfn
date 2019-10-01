package image

import (
	"fmt"
	"github.com/containers/image/transports"
	"github.com/containers/image/transports/alltransports"
	"github.com/containers/image/types"
	"github.com/slinkydeveloper/kfn/pkg/config"
	"strings"
)

type FunctionImage struct {
	ImageName string
	Tag       string
}

func (image FunctionImage) ParseSpecDest() (types.ImageReference, error) {
	destSpec := image.FullName()
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

func (image FunctionImage) FullName() string {
	if image.Tag != "" {
		return fmt.Sprintf("%s/%s:%s", config.ImageRegistry, image.ImageName, image.Tag)
	} else {
		return fmt.Sprintf("%s/%s", config.ImageRegistry, image.ImageName)
	}
}

func (image FunctionImage) FullNameForK8s() string {
	fullName := image.FullName()

	if strings.HasPrefix(fullName, "docker://") {
		return fullName[len("docker://"):]
	}

	return fullName
}
