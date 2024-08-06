package root

import (
	"github.com/dev.itbasis.sdkm/internal/version"
	"github.com/spf13/cobra"
)

var (
	CmdRoot = &cobra.Command{
		Version: version.BuildVersion(),
		Short:   "SDK Manager",
	}
)

func init() {
	CmdRoot.AddCommand(
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
