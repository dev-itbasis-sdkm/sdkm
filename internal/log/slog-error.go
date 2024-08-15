package log

import "log/slog"

func AttrError(err error) slog.Attr {
	return slog.Attr{Key: "error", Value: slog.StringValue(err.Error())}
}
