package os

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"slices"
	"strings"
)

// CleanPath Removes the sdkm directory from paths
//
//nolint:nonamedreturns // To optimize result logging
func CleanPath(path string, cleanPaths ...string) (result string) {
	var logger = slog.Default()

	defer slog.Debug("clean path: " + result)

	if logger.Enabled(context.Background(), slog.LevelDebug) {
		logger.Debug("path: " + path)
		logger.Debug(fmt.Sprintf("cleaning env path: %s", cleanPaths))
	}

	if len(cleanPaths) == 0 {
		result = path

		return result
	}

	var splitPaths = strings.Split(path, string(os.PathListSeparator))

	splitPaths = slices.DeleteFunc(
		splitPaths, func(s string) bool {
			return slices.Contains(cleanPaths, s)
		},
	)

	result = strings.Join(splitPaths, string(os.PathListSeparator))

	return result
}

func AddBeforePath(path string, paths ...string) string {
	if len(paths) == 0 {
		return path
	}

	return strings.Join(paths, string(os.PathListSeparator)) + string(os.PathListSeparator) + path
}
