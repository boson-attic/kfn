package js

import (
	"github.com/sirupsen/logrus"
	"github.com/slinkydeveloper/kfn/pkg"
	"github.com/slinkydeveloper/kfn/pkg/util"
	"io/ioutil"
	"os"
	"path"
)

const jsRuntimeRemoteZip = "https://github.com/openshift-cloud-functions/faas-js-runtime-image/archive/master.zip"

type jsRuntimeManager struct{}

func NewJSRuntimeManager() pkg.RuntimeManager {
	return jsRuntimeManager{}
}

func (j jsRuntimeManager) DownloadRuntimeIfRequired() error {
	if !util.FsExist(RuntimeDirectory()) {
		if err := util.MkdirpIfNotExists(RuntimeDirectory()); err != nil {
			return err
		}

		tempDir, err := ioutil.TempDir("", "faas-js-runtime-image")
		if err != nil {
			return err
		}

		runtimeZip := path.Join(tempDir, "master.zip")

		if err := util.DownloadFile(jsRuntimeRemoteZip, runtimeZip); err != nil {
			return err
		}

		logrus.Infof("Downloading runtime from %s to %s", jsRuntimeRemoteZip, runtimeZip)

		if _, err := util.Unzip(runtimeZip, tempDir); err != nil {
			return err
		}

		logrus.Infof("Runtime unzipped to %s", tempDir)

		if err := util.Copy(path.Join(tempDir, "faas-js-runtime-image-master", "src"), RuntimeDirectory()); err != nil {
			return err
		}
	} else {
		logrus.Infof("Using runtime cached in %s", RuntimeDirectory())
	}
	return nil
}

func (j jsRuntimeManager) ConfigureTargetDirectory(mainFile string, additionalFiles []string) error {
	err := os.Symlink(mainFile, path.Join(pkg.TargetDir, "usr", "index.js"))
	if err != nil {
		return err
	}

	for _, af := range additionalFiles {
		err := os.Symlink(af, path.Join(pkg.TargetDir, "usr", path.Base(af)))
		if err != nil {
			return err
		}
	}

	return os.Symlink(path.Join(RuntimeDirectory()), path.Join(pkg.TargetDir, "src"))
}

func RuntimeDirectory() string {
	return path.Join(pkg.RuntimeDir, "js")
}
