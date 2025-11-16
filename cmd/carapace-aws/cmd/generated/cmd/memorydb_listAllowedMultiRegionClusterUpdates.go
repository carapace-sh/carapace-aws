package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var memorydb_listAllowedMultiRegionClusterUpdatesCmd = &cobra.Command{
	Use:   "list-allowed-multi-region-cluster-updates",
	Short: "Lists the allowed updates for a multi-Region cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(memorydb_listAllowedMultiRegionClusterUpdatesCmd).Standalone()

	memorydb_listAllowedMultiRegionClusterUpdatesCmd.Flags().String("multi-region-cluster-name", "", "The name of the multi-Region cluster.")
	memorydb_listAllowedMultiRegionClusterUpdatesCmd.MarkFlagRequired("multi-region-cluster-name")
	memorydbCmd.AddCommand(memorydb_listAllowedMultiRegionClusterUpdatesCmd)
}
