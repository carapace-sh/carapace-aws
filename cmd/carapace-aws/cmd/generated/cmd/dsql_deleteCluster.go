package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dsql_deleteClusterCmd = &cobra.Command{
	Use:   "delete-cluster",
	Short: "Deletes a cluster in Amazon Aurora DSQL.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dsql_deleteClusterCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(dsql_deleteClusterCmd).Standalone()

		dsql_deleteClusterCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
		dsql_deleteClusterCmd.Flags().String("identifier", "", "The ID of the cluster to delete.")
		dsql_deleteClusterCmd.MarkFlagRequired("identifier")
	})
	dsqlCmd.AddCommand(dsql_deleteClusterCmd)
}
