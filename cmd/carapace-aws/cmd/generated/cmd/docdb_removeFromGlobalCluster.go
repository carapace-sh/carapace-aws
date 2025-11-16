package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var docdb_removeFromGlobalClusterCmd = &cobra.Command{
	Use:   "remove-from-global-cluster",
	Short: "Detaches an Amazon DocumentDB secondary cluster from a global cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(docdb_removeFromGlobalClusterCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(docdb_removeFromGlobalClusterCmd).Standalone()

		docdb_removeFromGlobalClusterCmd.Flags().String("db-cluster-identifier", "", "The Amazon Resource Name (ARN) identifying the cluster that was detached from the Amazon DocumentDB global cluster.")
		docdb_removeFromGlobalClusterCmd.Flags().String("global-cluster-identifier", "", "The cluster identifier to detach from the Amazon DocumentDB global cluster.")
		docdb_removeFromGlobalClusterCmd.MarkFlagRequired("db-cluster-identifier")
		docdb_removeFromGlobalClusterCmd.MarkFlagRequired("global-cluster-identifier")
	})
	docdbCmd.AddCommand(docdb_removeFromGlobalClusterCmd)
}
