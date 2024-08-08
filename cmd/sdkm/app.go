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
			prettylog.WithDestinationWriter(os.Stderr),
			prettylog.WithColor(),
		),
	)
	slog.SetDefault(logger)

	cmdRoot.SetOut(os.Stdout)
	cmdRoot.SetErr(os.Stderr)

	_ = cmdRoot.Execute()
}
