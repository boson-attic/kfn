package js

import (
	"strings"

	"github.com/pkg/errors"

	"github.com/slinkydeveloper/kfn/pkg/util"
)

type packageJsonConfig func(map[string]interface{})

func addLatestDevDependency(depName string) packageJsonConfig {
	return func(packageJson map[string]interface{}) {
		util.WriteNestedEntry(packageJson, depName, mustGetLatestNpmPackageVersionOrLatest(depName), "devDependencies")
	}
}

func addScript(scriptName string, scriptCommand string) packageJsonConfig {
	return func(packageJson map[string]interface{}) {
		scripts := packageJson["scripts"]
		if scripts == nil {
			scripts = make(map[string]string)
			packageJson["scripts"] = scripts
		}

		scripts.(map[string]string)[scriptName] = scriptCommand
	}
}

func getLatestNpmPackageVersion(pkgName string) (string, error) {
	v, err := util.RunCommandWithOutputBuffering("npm", []string{"show", pkgName, "version"}, "")
	if err != nil {
		return "", errors.Wrap(err, "error while trying to get npm package version")
	}

	return strings.Trim(v, "\n "), nil
}

func mustGetLatestNpmPackageVersionOrLatest(pkgName string) string {
	v, err := getLatestNpmPackageVersion(pkgName)
	if err != nil {
		return "latest"
	} else {
		return v
	}
}
