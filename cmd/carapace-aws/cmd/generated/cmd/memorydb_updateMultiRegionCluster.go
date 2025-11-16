package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var memorydb_updateMultiRegionClusterCmd = &cobra.Command{
	Use:   "update-multi-region-cluster",
	Short: "Updates the configuration of an existing multi-Region cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(memorydb_updateMultiRegionClusterCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(memorydb_updateMultiRegionClusterCmd).Standalone()

		memorydb_updateMultiRegionClusterCmd.Flags().String("description", "", "A new description for the multi-Region cluster.")
		memorydb_updateMultiRegionClusterCmd.Flags().String("engine-version", "", "The new engine version to be used for the multi-Region cluster.")
		memorydb_updateMultiRegionClusterCmd.Flags().String("multi-region-cluster-name", "", "The name of the multi-Region cluster to be updated.")
		memorydb_updateMultiRegionClusterCmd.Flags().String("multi-region-parameter-group-name", "", "The new multi-Region parameter group to be associated with the cluster.")
		memorydb_updateMultiRegionClusterCmd.Flags().String("node-type", "", "The new node type to be used for the multi-Region cluster.")
		memorydb_updateMultiRegionClusterCmd.Flags().String("shard-configuration", "", "")
		memorydb_updateMultiRegionClusterCmd.Flags().String("update-strategy", "", "The strategy to use for the update operation.")
		memorydb_updateMultiRegionClusterCmd.MarkFlagRequired("multi-region-cluster-name")
	})
	memorydbCmd.AddCommand(memorydb_updateMultiRegionClusterCmd)
}
