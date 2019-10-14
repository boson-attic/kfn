package util

import (
	"errors"
	"strings"
)

func CombineErrors(errs []error) error {
	s := make([]string, len(errs))
	for i, e := range errs {
		s[i] = e.Error()
	}
	return errors.New(strings.Join(s, "\n"))
}
