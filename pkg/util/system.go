package util

import (
	"bytes"
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

func RunCommandWithOutputBuffering(command string, params []string, workingDir string, envVariables ...[]string) (string, error) {
	var cmd = exec.Command(command, params...)

	cmd.Dir = workingDir

	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	cmd.Env = os.Environ()

	for _, envs := range envVariables {
		if envs != nil {
			cmd.Env = append(cmd.Env, envs...)
		}
	}

	err := cmd.Run()
	if err != nil {
		return "", err
	}

	if stderr.Len() != 0 {
		return "", errors.New(stderr.String())
	}

	return stdout.String(), nil
}
