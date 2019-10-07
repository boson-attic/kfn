package config

import (
	"io"
	"os"
)

func GetLoggerWriter() io.Writer {
	if Verbose || getEnvBoolOrDefault(DEBUG, false) {
		return os.Stdout
	} else {
		return NopLogger{}
	}
}

type NopLogger struct{}

func (n NopLogger) Write(p []byte) (int, error) {
	return len(p), nil
}
