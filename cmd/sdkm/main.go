package main

import (
	"log/slog"
	"os"

	"github.com/dev.itbasis.sdkm/cmd/sdkm/root"
	"github.com/dusted-go/logging/prettylog"
)

func main() {
	logger := slog.New(
		prettylog.New(
			&slog.HandlerOptions{Level: slog.LevelDebug},
			prettylog.WithDestinationWriter(os.Stdout),
			prettylog.WithColor(),
		),
	)
	slog.SetDefault(logger)

	_ = root.CmdRoot.Execute()
}
