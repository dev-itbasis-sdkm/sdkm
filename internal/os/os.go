package os

import (
	"fmt"
	"log"
	"log/slog"
	"os"
	"path/filepath"
	"strings"
)

const (
	DefaultDirMode  = 0o755
	DefaultFileMode = 0o644
)

func Pwd() string {
	executable, err := os.Getwd()
	if err != nil {
		log.Fatalln(err)
	}

	return executable
}

func UserHomeDir() string {
	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatalln(err)
	}

	return userHomeDir
}

// CleanEnvPath Removes the sdkm directory from paths
func CleanEnvPath(envPath string) string {
	executableDir := ExecutableDir()
	slog.Debug(fmt.Sprintf("cleaning env PATH: %s", executableDir))

	return strings.ReplaceAll(envPath, fmt.Sprintf("%s%c", executableDir, os.PathListSeparator), "")
}

func ExecutableDir() string {
	executable, err := os.Executable()
	if err != nil {
		log.Fatalln(err)
	}

	return filepath.Dir(executable)
}
