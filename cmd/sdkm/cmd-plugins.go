package sdkm

import (
	"strings"

	"github.com/dev.itbasis.sdkm/plugins"
	"github.com/spf13/cobra"
)

var cmdPlugins = &cobra.Command{
	Use:   "plugins",
	Short: "List plugins",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, _ []string) {
		cmd.Println("Available plugins: " + strings.Join(plugins.PluginNames, ", "))
	},
}
