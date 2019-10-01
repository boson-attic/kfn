package util

import (
	"github.com/gobuffalo/packr"
	"text/template"
)

type ResourceLoader struct {
	box packr.Box
}

func NewResourceLoader(path string) ResourceLoader {
	return ResourceLoader{box: packr.NewBox(path)}
}

func (rl ResourceLoader) LoadResource(filename string) ([]byte, error) {
	return rl.box.Find(filename)
}

func (rl ResourceLoader) LoadTemplate(filename string) (*template.Template, error) {
	templateStr, err := rl.box.FindString(filename)

	if err != nil {
		return nil, err
	}

	return template.New(filename).Parse(templateStr)
}
