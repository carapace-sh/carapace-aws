package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var neptune_failoverGlobalClusterCmd = &cobra.Command{
	Use:   "failover-global-cluster",
	Short: "Initiates the failover process for a Neptune global database.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(neptune_failoverGlobalClusterCmd).Standalone()

	neptune_failoverGlobalClusterCmd.Flags().String("allow-data-loss", "", "Specifies whether to allow data loss for this global database cluster operation.")
	neptune_failoverGlobalClusterCmd.Flags().String("global-cluster-identifier", "", "Identifier of the Neptune global database that should be failed over.")
	neptune_failoverGlobalClusterCmd.Flags().String("switchover", "", "Specifies whether to switch over this global database cluster.")
	neptune_failoverGlobalClusterCmd.Flags().String("target-db-cluster-identifier", "", "The Amazon Resource Name (ARN) of the secondary Neptune DB cluster that you want to promote to primary for the global database.")
	neptune_failoverGlobalClusterCmd.MarkFlagRequired("global-cluster-identifier")
	neptune_failoverGlobalClusterCmd.MarkFlagRequired("target-db-cluster-identifier")
	neptuneCmd.AddCommand(neptune_failoverGlobalClusterCmd)
}
