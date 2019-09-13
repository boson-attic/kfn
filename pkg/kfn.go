package pkg

import (
	"fmt"
	"github.com/slinkydeveloper/kfn/pkg/util"
	"os"
	"path"
	"strings"
)

const targetDirectoryBaseName = "target" //TODO use constant.go
const runtimeRemoteZip = "https://github.com/openshift-cloud-functions/faas-js-runtime-image/archive/master.zip"

var TargetDirectory string
var UsrDirectory string

func init() {
	// Init target directory
	pwd, err := os.Getwd()

	if err != nil {
		panic(fmt.Sprintf("Error while retrieving pwd: %v", err))
	}

	TargetDirectory = path.Join(pwd, targetDirectoryBaseName)
	UsrDirectory = path.Join(TargetDirectory, "usr")
}

func LoadFunction(location string) error {
	if err := initializeTargetDir(); err != nil {
		return err
	}

	if err := initializeUsrDir(); err != nil {
		return err
	}

	if strings.HasPrefix(location, "http") {
		return downloadFunctionFromHTTP(location)
	} else {
		return loadFunctionFileFromLocal(location)
	}
}

func DownloadRuntimeAndCopyRequiredFiles() error {
	if util.FsExist(path.Join(TargetDirectory, "src")) {
		// We reuse stuff already in target dir
		return nil
	}

	runtimeZip := path.Join(TargetDirectory, "master.zip")

	if err := util.DownloadFile(runtimeRemoteZip, runtimeZip); err != nil {
		return err
	}

	if _, err := util.Unzip(runtimeZip, TargetDirectory); err != nil {
		return err
	}

	srcDir := path.Join(TargetDirectory, "src")

	if err := util.Copy(path.Join(TargetDirectory, "faas-js-runtime-image-master", "src"), srcDir); err != nil {
		return err
	}

	return nil
}

func initializeTargetDir() error {
	if !util.FsExist(TargetDirectory) {
		return os.Mkdir(TargetDirectory, os.ModePerm)
	}
	return nil
}

func initializeUsrDir() error {
	if !util.FsExist(UsrDirectory) {
		return os.Mkdir(UsrDirectory, os.ModePerm)
	}
	return nil
}

func downloadFunctionFromHTTP(remote string) error {
	return util.DownloadFile(remote, path.Join(TargetDirectory, "usr", "fn.js"))
}

func loadFunctionFileFromLocal(filepath string) error {
	return util.Copy(filepath, path.Join(TargetDirectory, "usr", "index.js"))
}
