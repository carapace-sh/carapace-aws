package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var neptune_removeFromGlobalClusterCmd = &cobra.Command{
	Use:   "remove-from-global-cluster",
	Short: "Detaches a Neptune DB cluster from a Neptune global database.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(neptune_removeFromGlobalClusterCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(neptune_removeFromGlobalClusterCmd).Standalone()

		neptune_removeFromGlobalClusterCmd.Flags().String("db-cluster-identifier", "", "The Amazon Resource Name (ARN) identifying the cluster to be detached from the Neptune global database cluster.")
		neptune_removeFromGlobalClusterCmd.Flags().String("global-cluster-identifier", "", "The identifier of the Neptune global database from which to detach the specified Neptune DB cluster.")
		neptune_removeFromGlobalClusterCmd.MarkFlagRequired("db-cluster-identifier")
		neptune_removeFromGlobalClusterCmd.MarkFlagRequired("global-cluster-identifier")
	})
	neptuneCmd.AddCommand(neptune_removeFromGlobalClusterCmd)
}
