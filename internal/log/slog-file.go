package log

import "log/slog"

func AttrFilePath(filePath string) slog.Attr {
	return slog.String("filePath", filePath)
}
