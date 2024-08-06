package root

import (
	"os"
	"strings"

	sdkmOs "github.com/dev.itbasis.sdkm/internal/os"
	sdkmPlugins "github.com/dev.itbasis.sdkm/plugins"
	"github.com/spf13/cobra"
)

var cmdEnv = &cobra.Command{
	Use:        "env {" + strings.Join(sdkmPlugins.PluginNames, "|") + "} [<version>]",
	Short:      "Displays environment variables inside the environment used for the plugin",
	ArgAliases: []string{"plugin", "version"},
	Args:       cobra.MatchAll(cobra.RangeArgs(1, 2), cobra.OnlyValidArgs), //nolint:mnd //
	Run: func(cmd *cobra.Command, args []string) {
		getPluginFunc, ok := sdkmPlugins.Plugins[args[0]]
		if !ok {
			cmd.PrintErrf("plugin %s not found", args[0])
			os.Exit(1)
		}

		sdkmPlugin := getPluginFunc()

		//nolint:mnd //
		if len(args) >= 2 {
			for k, v := range sdkmPlugin.EnvByVersion(cmd.Context(), args[1]) {
				cmd.Printf("%s=%s\n", k, v)
			}
		} else {
			environ, errEnv := sdkmPlugin.Env(cmd.Context(), sdkmOs.Pwd())
			if errEnv != nil {
				cmd.PrintErrln(errEnv)
				os.Exit(1)
			}

			for k, v := range environ {
				cmd.Printf("%s=%s\n", k, v)
			}
		}
	},
}
