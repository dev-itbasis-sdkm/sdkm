package sdkm

import (
	"log/slog"
	"os"

	"github.com/dusted-go/logging/prettylog"
)

var (
	logLevel = &slog.LevelVar{}

	logger = slog.New(
		prettylog.New(
			&slog.HandlerOptions{Level: logLevel},
			prettylog.WithDestinationWriter(os.Stderr),
			prettylog.WithColor(),
		),
	)
)
