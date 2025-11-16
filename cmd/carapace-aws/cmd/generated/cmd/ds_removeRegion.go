package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ds_removeRegionCmd = &cobra.Command{
	Use:   "remove-region",
	Short: "Stops all replication and removes the domain controllers from the specified Region.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ds_removeRegionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ds_removeRegionCmd).Standalone()

		ds_removeRegionCmd.Flags().String("directory-id", "", "The identifier of the directory for which you want to remove Region replication.")
		ds_removeRegionCmd.MarkFlagRequired("directory-id")
	})
	dsCmd.AddCommand(ds_removeRegionCmd)
}
