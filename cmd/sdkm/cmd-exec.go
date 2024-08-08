package sdkm

import (
	"os"
	"strings"

	sdkmOs "github.com/dev.itbasis.sdkm/internal/os"
	sdkmPlugins "github.com/dev.itbasis.sdkm/plugins"
	"github.com/spf13/cobra"
)

var cmdExec = &cobra.Command{
	Use:                "exec {" + strings.Join(sdkmPlugins.PluginNames, "|") + "} {<program>} [<args...>]",
	Short:              "Execute a command in a plugin",
	DisableFlagParsing: true,
	Args:               cobra.MatchAll(cobra.MinimumNArgs(2)), //nolint:mnd // TODO
	ArgAliases:         []string{"plugin", "program"},
	RunE: func(cmd *cobra.Command, args []string) error {
		getPluginFunc, ok := sdkmPlugins.Plugins[args[0]]
		if !ok {
			cmd.PrintErrf("plugin %s not found", args[0])
			os.Exit(1)
		}

		//nolint:wrapcheck // TODO
		return getPluginFunc().
			Exec(cmd.Context(), sdkmOs.Pwd(), cmd.InOrStdin(), cmd.OutOrStdout(), cmd.OutOrStderr(), args[1:])
	},
}
