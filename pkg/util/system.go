package util

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/pkg/errors"

	"github.com/slinkydeveloper/kfn/pkg/config"
)

func CommandsExists(cmd ...string) error {
	for _, c := range cmd {
		if _, err := exec.LookPath(c); err != nil {
			return errors.Wrap(err, fmt.Sprintf("cannot find `%s`", c))
		}
	}
	return nil
}

func RunCommand(command string, params []string, workingDir string, envVariables ...[]string) error {
	var cmd = exec.Command(command, params...)

	cmd.Dir = workingDir
	// Configure proper logging
	cmd.Stdout = config.GetLoggerWriter()
	cmd.Stderr = config.GetLoggerWriter()

	cmd.Env = os.Environ()

	for _, envs := range envVariables {
		if envs != nil {
			cmd.Env = append(cmd.Env, envs...)
		}
	}

	return cmd.Run()
}
