package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var memorydb_deleteMultiRegionClusterCmd = &cobra.Command{
	Use:   "delete-multi-region-cluster",
	Short: "Deletes an existing multi-Region cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(memorydb_deleteMultiRegionClusterCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(memorydb_deleteMultiRegionClusterCmd).Standalone()

		memorydb_deleteMultiRegionClusterCmd.Flags().String("multi-region-cluster-name", "", "The name of the multi-Region cluster to be deleted.")
		memorydb_deleteMultiRegionClusterCmd.MarkFlagRequired("multi-region-cluster-name")
	})
	memorydbCmd.AddCommand(memorydb_deleteMultiRegionClusterCmd)
}
