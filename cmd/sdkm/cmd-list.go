package sdkm

import (
	"github.com/spf13/cobra"
)

var cmdList = &cobra.Command{
	Use:   "list",
	Short: "List installed versions",
}

func init() {
	cmdList.AddCommand(CmdListAll)
}
