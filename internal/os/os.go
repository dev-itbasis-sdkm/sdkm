package os

import (
	"log"
	"os"
	"path/filepath"
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

func ExecutableDir() string {
	executable, err := os.Executable()
	if err != nil {
		log.Fatalln(err)
	}

	return filepath.Dir(executable)
}
