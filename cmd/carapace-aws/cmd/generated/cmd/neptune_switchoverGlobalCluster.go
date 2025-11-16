package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var neptune_switchoverGlobalClusterCmd = &cobra.Command{
	Use:   "switchover-global-cluster",
	Short: "Switches over the specified secondary DB cluster to be the new primary DB cluster in the global database cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(neptune_switchoverGlobalClusterCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(neptune_switchoverGlobalClusterCmd).Standalone()

		neptune_switchoverGlobalClusterCmd.Flags().String("global-cluster-identifier", "", "The identifier of the global database cluster to switch over.")
		neptune_switchoverGlobalClusterCmd.Flags().String("target-db-cluster-identifier", "", "The Amazon Resource Name (ARN) of the secondary Neptune DB cluster that you want to promote to primary for the global database.")
		neptune_switchoverGlobalClusterCmd.MarkFlagRequired("global-cluster-identifier")
		neptune_switchoverGlobalClusterCmd.MarkFlagRequired("target-db-cluster-identifier")
	})
	neptuneCmd.AddCommand(neptune_switchoverGlobalClusterCmd)
}
