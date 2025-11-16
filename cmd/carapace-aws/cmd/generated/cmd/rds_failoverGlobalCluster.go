package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rds_failoverGlobalClusterCmd = &cobra.Command{
	Use:   "failover-global-cluster",
	Short: "Promotes the specified secondary DB cluster to be the primary DB cluster in the global database cluster to fail over or switch over a global database.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rds_failoverGlobalClusterCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(rds_failoverGlobalClusterCmd).Standalone()

		rds_failoverGlobalClusterCmd.Flags().String("allow-data-loss", "", "Specifies whether to allow data loss for this global database cluster operation.")
		rds_failoverGlobalClusterCmd.Flags().String("global-cluster-identifier", "", "The identifier of the global database cluster (Aurora global database) this operation should apply to.")
		rds_failoverGlobalClusterCmd.Flags().String("switchover", "", "Specifies whether to switch over this global database cluster.")
		rds_failoverGlobalClusterCmd.Flags().String("target-db-cluster-identifier", "", "The identifier of the secondary Aurora DB cluster that you want to promote to the primary for the global database cluster.")
		rds_failoverGlobalClusterCmd.MarkFlagRequired("global-cluster-identifier")
		rds_failoverGlobalClusterCmd.MarkFlagRequired("target-db-cluster-identifier")
	})
	rdsCmd.AddCommand(rds_failoverGlobalClusterCmd)
}
