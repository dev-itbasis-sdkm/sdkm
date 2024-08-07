package sdkm

import (
	"log/slog"
	"os"

	"github.com/dusted-go/logging/prettylog"
)

type App struct {
}

func NewApp() *App {
	return &App{}
}

func (receiver *App) Run() {
	logger := slog.New(
		prettylog.New(
			&slog.HandlerOptions{Level: slog.LevelDebug},
			prettylog.WithDestinationWriter(os.Stdout),
			prettylog.WithColor(),
		),
	)
	slog.SetDefault(logger)

	_ = cmdRoot.Execute()
}
