package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var finspace_deleteKxClusterCmd = &cobra.Command{
	Use:   "delete-kx-cluster",
	Short: "Deletes a kdb cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(finspace_deleteKxClusterCmd).Standalone()

	finspace_deleteKxClusterCmd.Flags().String("client-token", "", "A token that ensures idempotency.")
	finspace_deleteKxClusterCmd.Flags().String("cluster-name", "", "The name of the cluster that you want to delete.")
	finspace_deleteKxClusterCmd.Flags().String("environment-id", "", "A unique identifier for the kdb environment.")
	finspace_deleteKxClusterCmd.MarkFlagRequired("cluster-name")
	finspace_deleteKxClusterCmd.MarkFlagRequired("environment-id")
	finspaceCmd.AddCommand(finspace_deleteKxClusterCmd)
}
