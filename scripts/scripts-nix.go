//go:build !windows

package scripts

import (
	"errors"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
)

func Unpack(targetDir string) error {
	dirEntries, errReadDir := scripts.ReadDir(".")
	if errReadDir != nil {
		return errReadDir
	}

	var errUnpack error

	for _, dirEntry := range dirEntries {
		entryName := dirEntry.Name()

		bytes, errReadFile := scripts.ReadFile(entryName)
		if errReadFile != nil {
			errUnpack = errors.Join(errUnpack, errReadFile)
		}

		slog.Info(fmt.Sprintf("Unpacking file: %s", entryName))

		var fileMode os.FileMode = 0744
		// if strings.Contains(entryName, ".") {
		// 	fileMode = 0644
		// }

		if errWrite := os.WriteFile(filepath.Join(targetDir, entryName), bytes, fileMode); errWrite != nil {
			errUnpack = errors.Join(errUnpack, errWrite)
		}
	}

	return nil
}
