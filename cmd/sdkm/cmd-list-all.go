package sdkm

import (
	"log"
	"strings"

	sdkmSDKVersion "github.com/dev.itbasis.sdkm/pkg/sdk-version"
	sdkmPlugins "github.com/dev.itbasis.sdkm/plugins"
	"github.com/spf13/cobra"
)

var CmdListAll = &cobra.Command{
	Use:        "all {" + strings.Join(sdkmPlugins.PluginNames, "|") + "} [<version>]",
	Short:      "List all versions",
	ArgAliases: []string{"plugin", "version"},
	Args:       cobra.MatchAll(cobra.RangeArgs(1, 2), cobra.OnlyValidArgs), //nolint:mnd // TODO
	Run: func(cmd *cobra.Command, args []string) {
		getPluginFunc, ok := sdkmPlugins.Plugins[args[0]]
		if !ok {
			log.Fatalf("plugin %s not found", args[0])
		}

		var sdkVersions []sdkmSDKVersion.SDKVersion

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
