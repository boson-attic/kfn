package image

import "fmt"

type FunctionImage struct {
	ImageName     string
	ImageRegistry string
	Tag           string
}

func (image FunctionImage) FullName() string {
	return fmt.Sprintf("%s/%s:%s", image.ImageRegistry, image.ImageName, image.Tag)
}
