package root

import (
	"os"
	"strings"

	sdkmOs "github.com/dev.itbasis.sdkm/internal/os"
	sdkmPlugins "github.com/dev.itbasis.sdkm/plugins"
	"github.com/spf13/cobra"
)

var cmdCurrent = &cobra.Command{
	Use:        "current [" + strings.Join(sdkmPlugins.PluginNames, "|") + "]",
	Short:      "Display current version",
	ArgAliases: []string{"plugin"},
	Args:       cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
	Run: func(cmd *cobra.Command, args []string) {
		getPluginFunc, ok := sdkmPlugins.Plugins[args[0]]
		if !ok {
			cmd.PrintErrf("plugin %s not found", args[0])
			os.Exit(1)
		}

		sdkVersion, err := getPluginFunc().Current(cmd.Context(), sdkmOs.Pwd())
		if err != nil {
			cmd.PrintErr(err)
			os.Exit(1)
		}

		cmd.Print(sdkVersion.ID)
		if !sdkVersion.Installed {
			cmd.Print(" [not installed]")
		}
		cmd.Println()
	},
}
