package editors

import (
	"fmt"
	"os/exec"

	log "github.com/sirupsen/logrus"
	"github.com/slinkydeveloper/kfn/pkg/util"
)

type Editor uint8

const (
	Idea Editor = iota
	VSCode
	Unknown
)

func GetEditor(editor string) Editor {
	switch editor {
	case "idea":
		return Idea
	case "vscode":
		return VSCode
	case "code":
		return VSCode
	}
	return Unknown
}

func (e Editor) OpenEditor(directory string) error {
	switch e {
	case Idea:
		return startEditorProcess(directory, "idea", directory)
	case VSCode:
		return startEditorProcess(directory, "code", directory)
	}
	return fmt.Errorf("Editor not supported")
}

func startEditorProcess(pwd string, processName string, args ...string) error {
	if err := util.CommandsExists(processName); err != nil {
		return err
	}
	log.Infof("Starting process %s with args %v and pwd %s", processName, args, pwd)
	cmd := exec.Command(processName, args...)
	cmd.Dir = pwd
	return cmd.Start()
}
