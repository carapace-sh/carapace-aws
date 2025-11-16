package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var memorydb_deleteClusterCmd = &cobra.Command{
	Use:   "delete-cluster",
	Short: "Deletes a cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(memorydb_deleteClusterCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(memorydb_deleteClusterCmd).Standalone()

		memorydb_deleteClusterCmd.Flags().String("cluster-name", "", "The name of the cluster to be deleted")
		memorydb_deleteClusterCmd.Flags().String("final-snapshot-name", "", "The user-supplied name of a final cluster snapshot.")
		memorydb_deleteClusterCmd.Flags().String("multi-region-cluster-name", "", "The name of the multi-Region cluster to be deleted.")
		memorydb_deleteClusterCmd.MarkFlagRequired("cluster-name")
	})
	memorydbCmd.AddCommand(memorydb_deleteClusterCmd)
}
