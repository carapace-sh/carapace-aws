package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ds_getDirectoryLimitsCmd = &cobra.Command{
	Use:   "get-directory-limits",
	Short: "Obtains directory limit information for the current Region.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ds_getDirectoryLimitsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ds_getDirectoryLimitsCmd).Standalone()

	})
	dsCmd.AddCommand(ds_getDirectoryLimitsCmd)
}
