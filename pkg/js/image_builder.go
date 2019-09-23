package js

import (
	"context"
	"github.com/containers/image/types"
	"github.com/slinkydeveloper/kfn/pkg"
	"github.com/slinkydeveloper/kfn/pkg/image"
	"path"
)

const (
	baseImage = "node:12-alpine"
)

type jsImageBuilder struct{}

func NewJSImageBuilder() pkg.ImageBuilder {
	return jsImageBuilder{}
}

func (j jsImageBuilder) BuildImage(systemContext *types.SystemContext, imageName string, imageTag string) (image.FunctionImage, error) {
	builder, err := pkg.InitializeBuilder(context.TODO(), systemContext, baseImage)
	if err != nil {
		return image.FunctionImage{}, err
	}

	builder.SetPort("8080")

	err = pkg.Add(builder, pkg.BuildAdd{From: path.Join(pkg.TargetDir, "src"), To: "/home/node/src"}, pkg.BuildAdd{From: path.Join(pkg.TargetDir, "usr"), To: "/home/node/usr"})
	if err != nil {
		return image.FunctionImage{}, err
	}

	err = pkg.RunCommands(
		builder,
		pkg.BuildCommand{Command: "mkdir -p /home/node/usr/.npm"},
		pkg.BuildCommand{Command: "chmod -R a+g+x /home/node/usr"},
		pkg.BuildCommand{Command: "chmod -R a+g+x /home/node/src"},
		pkg.BuildCommand{"npm install", "/home/node/usr"},
		pkg.BuildCommand{"npm install", "/home/node/src"},
	)
	if err != nil {
		return image.FunctionImage{}, err
	}

	builder.SetEnv("HOME", "/home/node/usr")
	builder.SetUser("1000")
	builder.SetWorkDir("/home/node/src")

	builder.SetCmd([]string{"node", "/home/node/src/index.js"})

	return pkg.CommitImage(builder, imageName, imageTag)
}
