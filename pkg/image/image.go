package image

import (
	"fmt"
	"github.com/slinkydeveloper/kfn/pkg"
	"strings"
)

type FunctionImage struct {
	ImageId   string
	ImageName string
	Tag       string
}

func (image FunctionImage) FullName() string {
	if image.Tag != "" {
		return fmt.Sprintf("%s/%s:%s", pkg.ImageRegistry, image.ImageName, image.Tag)
	} else {
		return fmt.Sprintf("%s/%s", pkg.ImageRegistry, image.ImageName)
	}
}

func (image FunctionImage) FullNameForK8s() string {
	fullName := image.FullName()

	if strings.HasPrefix(fullName, "docker://") {
		return fullName[len("docker://"):]
	}

	return fullName
}
