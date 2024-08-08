package sdkm

import (
	"log/slog"
	"os"
)

type App struct {
}

func NewApp() *App {
	return &App{}
}

func (receiver *App) Run() {
	logLevel.Set(slog.LevelInfo)

	slog.SetDefault(logger)

	cmdRoot.SetOut(os.Stdout)
	cmdRoot.SetErr(os.Stderr)

	_ = cmdRoot.Execute()
}
