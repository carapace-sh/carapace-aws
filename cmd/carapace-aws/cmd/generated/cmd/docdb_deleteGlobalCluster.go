package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var docdb_deleteGlobalClusterCmd = &cobra.Command{
	Use:   "delete-global-cluster",
	Short: "Deletes a global cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(docdb_deleteGlobalClusterCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(docdb_deleteGlobalClusterCmd).Standalone()

		docdb_deleteGlobalClusterCmd.Flags().String("global-cluster-identifier", "", "The cluster identifier of the global cluster being deleted.")
		docdb_deleteGlobalClusterCmd.MarkFlagRequired("global-cluster-identifier")
	})
	docdbCmd.AddCommand(docdb_deleteGlobalClusterCmd)
}
