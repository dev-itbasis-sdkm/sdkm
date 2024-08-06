package root

import (
	sdkmOs "github.com/dev.itbasis.sdkm/internal/os"
	"github.com/dev.itbasis.sdkm/scripts"
	"github.com/spf13/cobra"
)

var cmdReshim = &cobra.Command{
	Use:   "reshim",
	Short: "Unpacking scripts and shims",
	RunE: func(_ *cobra.Command, _ []string) error {
		return scripts.Unpack(sdkmOs.ExecutableDir())
	},
}
