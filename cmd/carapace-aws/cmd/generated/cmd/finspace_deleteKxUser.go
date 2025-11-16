package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var finspace_deleteKxUserCmd = &cobra.Command{
	Use:   "delete-kx-user",
	Short: "Deletes a user in the specified kdb environment.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(finspace_deleteKxUserCmd).Standalone()

	finspace_deleteKxUserCmd.Flags().String("client-token", "", "A token that ensures idempotency.")
	finspace_deleteKxUserCmd.Flags().String("environment-id", "", "A unique identifier for the kdb environment.")
	finspace_deleteKxUserCmd.Flags().String("user-name", "", "A unique identifier for the user that you want to delete.")
	finspace_deleteKxUserCmd.MarkFlagRequired("environment-id")
	finspace_deleteKxUserCmd.MarkFlagRequired("user-name")
	finspaceCmd.AddCommand(finspace_deleteKxUserCmd)
}
