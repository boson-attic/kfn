package image

import (
	"fmt"
	"strings"
)

type FunctionImage struct {
	ImageName     string
	ImageRegistry string
	Tag           string
}

func (image FunctionImage) FullName() string {
	if image.Tag != "" {
		return fmt.Sprintf("%s/%s:%s", image.ImageRegistry, image.ImageName, image.Tag)
	} else {
		return fmt.Sprintf("%s/%s", image.ImageRegistry, image.ImageName)
	}
}

func (image FunctionImage) FullNameForK8s() string {
	fullName := image.FullName()

	if strings.HasPrefix(fullName, "docker://") {
		return fullName[len("docker://"):]
	}

	return fullName
}
