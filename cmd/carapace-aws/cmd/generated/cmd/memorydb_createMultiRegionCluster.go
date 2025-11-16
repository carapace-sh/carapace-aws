package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var memorydb_createMultiRegionClusterCmd = &cobra.Command{
	Use:   "create-multi-region-cluster",
	Short: "Creates a new multi-Region cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(memorydb_createMultiRegionClusterCmd).Standalone()

	memorydb_createMultiRegionClusterCmd.Flags().String("description", "", "A description for the multi-Region cluster.")
	memorydb_createMultiRegionClusterCmd.Flags().String("engine", "", "The name of the engine to be used for the multi-Region cluster.")
	memorydb_createMultiRegionClusterCmd.Flags().String("engine-version", "", "The version of the engine to be used for the multi-Region cluster.")
	memorydb_createMultiRegionClusterCmd.Flags().String("multi-region-cluster-name-suffix", "", "A suffix to be added to the Multi-Region cluster name.")
	memorydb_createMultiRegionClusterCmd.Flags().String("multi-region-parameter-group-name", "", "The name of the multi-Region parameter group to be associated with the cluster.")
	memorydb_createMultiRegionClusterCmd.Flags().String("node-type", "", "The node type to be used for the multi-Region cluster.")
	memorydb_createMultiRegionClusterCmd.Flags().String("num-shards", "", "The number of shards for the multi-Region cluster.")
	memorydb_createMultiRegionClusterCmd.Flags().String("tags", "", "A list of tags to be applied to the multi-Region cluster.")
	memorydb_createMultiRegionClusterCmd.Flags().String("tlsenabled", "", "Whether to enable TLS encryption for the multi-Region cluster.")
	memorydb_createMultiRegionClusterCmd.MarkFlagRequired("multi-region-cluster-name-suffix")
	memorydb_createMultiRegionClusterCmd.MarkFlagRequired("node-type")
	memorydbCmd.AddCommand(memorydb_createMultiRegionClusterCmd)
}
