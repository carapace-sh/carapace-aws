package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var memorydb_describeMultiRegionClustersCmd = &cobra.Command{
	Use:   "describe-multi-region-clusters",
	Short: "Returns details about one or more multi-Region clusters.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(memorydb_describeMultiRegionClustersCmd).Standalone()

	memorydb_describeMultiRegionClustersCmd.Flags().String("max-results", "", "The maximum number of results to return.")
	memorydb_describeMultiRegionClustersCmd.Flags().String("multi-region-cluster-name", "", "The name of a specific multi-Region cluster to describe.")
	memorydb_describeMultiRegionClustersCmd.Flags().String("next-token", "", "A token to specify where to start paginating.")
	memorydb_describeMultiRegionClustersCmd.Flags().String("show-cluster-details", "", "Details about the multi-Region cluster.")
	memorydbCmd.AddCommand(memorydb_describeMultiRegionClustersCmd)
}
