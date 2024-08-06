package root

import (
	"log"
	"strings"

	"github.com/dev.itbasis.sdkm/pkg/plugin"
	"github.com/dev.itbasis.sdkm/plugins"
	"github.com/spf13/cobra"
)

var CmdListAll = &cobra.Command{
	Use:        "all {" + strings.Join(plugins.PluginNames, "|") + "} [<version>]",
	Short:      "List all versions",
	ArgAliases: []string{"plugin", "version"},
	Args:       cobra.MatchAll(cobra.RangeArgs(1, 2), cobra.OnlyValidArgs), //nolint:mnd //
	Run: func(cmd *cobra.Command, args []string) {
		getPluginFunc, ok := plugins.Plugins[args[0]]
		if !ok {
			log.Fatalf("plugin %s not found", args[0])
		}

		var sdkVersions []plugin.SDKVersion

		if len(args) == 1 {
			sdkVersions = getPluginFunc().ListAllVersions(cmd.Context())
		} else {
			sdkVersions = getPluginFunc().ListAllVersionsByPrefix(cmd.Context(), args[1])
		}

		for _, version := range sdkVersions {
			cmd.Println(version)
		}
	},
}
