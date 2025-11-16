package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rds_removeFromGlobalClusterCmd = &cobra.Command{
	Use:   "remove-from-global-cluster",
	Short: "Detaches an Aurora secondary cluster from an Aurora global database cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rds_removeFromGlobalClusterCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(rds_removeFromGlobalClusterCmd).Standalone()

		rds_removeFromGlobalClusterCmd.Flags().String("db-cluster-identifier", "", "The Amazon Resource Name (ARN) identifying the cluster that was detached from the Aurora global database cluster.")
		rds_removeFromGlobalClusterCmd.Flags().String("global-cluster-identifier", "", "The cluster identifier to detach from the Aurora global database cluster.")
		rds_removeFromGlobalClusterCmd.MarkFlagRequired("db-cluster-identifier")
		rds_removeFromGlobalClusterCmd.MarkFlagRequired("global-cluster-identifier")
	})
	rdsCmd.AddCommand(rds_removeFromGlobalClusterCmd)
}
