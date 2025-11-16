package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var memorydb_batchUpdateClusterCmd = &cobra.Command{
	Use:   "batch-update-cluster",
	Short: "Apply the service update to a list of clusters supplied.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(memorydb_batchUpdateClusterCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(memorydb_batchUpdateClusterCmd).Standalone()

		memorydb_batchUpdateClusterCmd.Flags().String("cluster-names", "", "The cluster names to apply the updates.")
		memorydb_batchUpdateClusterCmd.Flags().String("service-update", "", "The unique ID of the service update")
		memorydb_batchUpdateClusterCmd.MarkFlagRequired("cluster-names")
	})
	memorydbCmd.AddCommand(memorydb_batchUpdateClusterCmd)
}
