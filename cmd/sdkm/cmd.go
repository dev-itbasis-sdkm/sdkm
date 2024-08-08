package sdkm

import (
	"log/slog"

	"github.com/dev.itbasis.sdkm/internal/version"
	"github.com/spf13/cobra"
)

var (
	cmdRoot = &cobra.Command{
		Version: version.BuildVersion(),
		Short:   "SDK Manager",
	}

	debugFlag = cmdRoot.PersistentFlags().BoolP("debug", "d", false, "debug mode")
)

func init() {
	cmdRoot.PersistentPreRun = func(_ *cobra.Command, _ []string) {
		if *debugFlag {
			logLevel.Set(slog.LevelDebug)
		}
	}

	cmdRoot.AddCommand(
		cmdPlugins,
		cmdList,
		cmdLatest,
		cmdCurrent,
		cmdInstall,
		cmdEnv,
		cmdExec,
		cmdReshim,
	)
}
