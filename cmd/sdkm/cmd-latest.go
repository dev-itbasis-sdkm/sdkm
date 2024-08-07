package sdkm

import (
	"log"
	"os"
	"strings"

	"github.com/dev.itbasis.sdkm/plugins"
	"github.com/spf13/cobra"
)

var cmdLatest = &cobra.Command{
	Use:        "latest {" + strings.Join(plugins.PluginNames, "|") + "} [<version>]",
	Short:      "Show latest stable version",
	ArgAliases: []string{"plugin", "version"},
	Args:       cobra.MatchAll(cobra.RangeArgs(1, 2), cobra.OnlyValidArgs), //nolint:mnd // TODO
	Run: func(cmd *cobra.Command, args []string) {
		getPluginFunc, ok := plugins.Plugins[args[0]]
		if !ok {
			log.Fatalf("plugin %s not found", args[0])
		}

		sdkmPlugin := getPluginFunc()

		if len(args) == 1 {
			sdkVersion := sdkmPlugin.LatestVersion(cmd.Context())
			cmd.Println(sdkVersion.Print())

			os.Exit(0)
		}

		sdkVersion, err := sdkmPlugin.LatestVersionByPrefix(cmd.Context(), args[1])
		if err != nil {
			cmd.PrintErrln(err)
			os.Exit(1)
		}

		cmd.Println(sdkVersion.PrintWithOptions(false, true, true))
	},
}
