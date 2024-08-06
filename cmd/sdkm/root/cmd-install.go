package root

import (
	"os"
	"strings"

	sdkmOs "github.com/dev.itbasis.sdkm/internal/os"
	sdkmPlugins "github.com/dev.itbasis.sdkm/plugins"
	"github.com/spf13/cobra"
)

var cmdInstall = &cobra.Command{
	Use:        "install [" + strings.Join(sdkmPlugins.PluginNames, "|") + "] [<version>]",
	Short:      "Install the SDK",
	Args:       cobra.MatchAll(cobra.RangeArgs(1, 2), cobra.OnlyValidArgs), //nolint:mnd //
	ArgAliases: []string{"plugin", "version"},
	RunE: func(cmd *cobra.Command, args []string) error {
		getPluginFunc, ok := sdkmPlugins.Plugins[args[0]]
		if !ok {
			cmd.PrintErrf("plugin %s not found", args[0])
			os.Exit(1)
		}

		sdkmPlugin := getPluginFunc()

		//nolint:mnd //
		if len(args) == 2 {
			return sdkmPlugin.InstallVersion(cmd.Context(), args[1]) //nolint:wrapcheck
		}

		return sdkmPlugin.Install(cmd.Context(), sdkmOs.Pwd()) //nolint:wrapcheck
	},
}
