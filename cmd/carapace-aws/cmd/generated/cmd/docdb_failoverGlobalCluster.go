package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var docdb_failoverGlobalClusterCmd = &cobra.Command{
	Use:   "failover-global-cluster",
	Short: "Promotes the specified secondary DB cluster to be the primary DB cluster in the global cluster when failing over a global cluster occurs.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(docdb_failoverGlobalClusterCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(docdb_failoverGlobalClusterCmd).Standalone()

		docdb_failoverGlobalClusterCmd.Flags().String("allow-data-loss", "", "Specifies whether to allow data loss for this global cluster operation.")
		docdb_failoverGlobalClusterCmd.Flags().String("global-cluster-identifier", "", "The identifier of the Amazon DocumentDB global cluster to apply this operation.")
		docdb_failoverGlobalClusterCmd.Flags().String("switchover", "", "Specifies whether to switch over this global database cluster.")
		docdb_failoverGlobalClusterCmd.Flags().String("target-db-cluster-identifier", "", "The identifier of the secondary Amazon DocumentDB cluster that you want to promote to the primary for the global cluster.")
		docdb_failoverGlobalClusterCmd.MarkFlagRequired("global-cluster-identifier")
		docdb_failoverGlobalClusterCmd.MarkFlagRequired("target-db-cluster-identifier")
	})
	docdbCmd.AddCommand(docdb_failoverGlobalClusterCmd)
}
