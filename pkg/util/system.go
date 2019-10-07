package util

import (
	"fmt"
	"github.com/pkg/errors"
	"os/exec"
)

func CommandsExists(cmd ...string) error {
	for _, c := range cmd {
		if _, err := exec.LookPath(c); err != nil {
			return errors.Wrap(err, fmt.Sprintf("cannot find `%s`", c))
		}
	}
	return nil
}
