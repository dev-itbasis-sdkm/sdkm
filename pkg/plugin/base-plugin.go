package plugin

import (
	"io"
)

//go:generate mockgen -source=$GOFILE -package=$GOPACKAGE -destination=base-plugin.mock.go
type BasePlugin interface {
	WithSDKDir(dir string) BasePlugin

	GetSDKDir() string
	GetSDKVersionDir(pluginName, version string) string
	HasInstalled(pluginName, version string) bool

	Exec(overrideEnv map[string]string, stdIn io.Reader, stdOut, stdErr io.Writer, args []string) error
}
