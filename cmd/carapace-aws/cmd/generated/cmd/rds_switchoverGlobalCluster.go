package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rds_switchoverGlobalClusterCmd = &cobra.Command{
	Use:   "switchover-global-cluster",
	Short: "Switches over the specified secondary DB cluster to be the new primary DB cluster in the global database cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rds_switchoverGlobalClusterCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(rds_switchoverGlobalClusterCmd).Standalone()

		rds_switchoverGlobalClusterCmd.Flags().String("global-cluster-identifier", "", "The identifier of the global database cluster to switch over.")
		rds_switchoverGlobalClusterCmd.Flags().String("target-db-cluster-identifier", "", "The identifier of the secondary Aurora DB cluster to promote to the new primary for the global database cluster.")
		rds_switchoverGlobalClusterCmd.MarkFlagRequired("global-cluster-identifier")
		rds_switchoverGlobalClusterCmd.MarkFlagRequired("target-db-cluster-identifier")
	})
	rdsCmd.AddCommand(rds_switchoverGlobalClusterCmd)
}
