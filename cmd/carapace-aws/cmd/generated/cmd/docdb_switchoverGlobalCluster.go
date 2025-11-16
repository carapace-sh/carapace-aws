package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var docdb_switchoverGlobalClusterCmd = &cobra.Command{
	Use:   "switchover-global-cluster",
	Short: "Switches over the specified secondary Amazon DocumentDB cluster to be the new primary Amazon DocumentDB cluster in the global database cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(docdb_switchoverGlobalClusterCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(docdb_switchoverGlobalClusterCmd).Standalone()

		docdb_switchoverGlobalClusterCmd.Flags().String("global-cluster-identifier", "", "The identifier of the Amazon DocumentDB global database cluster to switch over.")
		docdb_switchoverGlobalClusterCmd.Flags().String("target-db-cluster-identifier", "", "The identifier of the secondary Amazon DocumentDB cluster to promote to the new primary for the global database cluster.")
		docdb_switchoverGlobalClusterCmd.MarkFlagRequired("global-cluster-identifier")
		docdb_switchoverGlobalClusterCmd.MarkFlagRequired("target-db-cluster-identifier")
	})
	docdbCmd.AddCommand(docdb_switchoverGlobalClusterCmd)
}
