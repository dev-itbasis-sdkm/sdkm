package sdkm

import (
	"github.com/dev.itbasis.sdkm/internal/version"
	"github.com/spf13/cobra"
)

var (
	cmdRoot = &cobra.Command{
		Version: version.BuildVersion(),
		Short:   "SDK Manager",
	}
)

func init() {
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
