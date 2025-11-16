package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pcs_deleteClusterCmd = &cobra.Command{
	Use:   "delete-cluster",
	Short: "Deletes a cluster and all its linked resources.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pcs_deleteClusterCmd).Standalone()

	pcs_deleteClusterCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
	pcs_deleteClusterCmd.Flags().String("cluster-identifier", "", "The name or ID of the cluster to delete.")
	pcs_deleteClusterCmd.MarkFlagRequired("cluster-identifier")
	pcsCmd.AddCommand(pcs_deleteClusterCmd)
}
