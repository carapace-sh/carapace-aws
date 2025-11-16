package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var finspace_deleteKxEnvironmentCmd = &cobra.Command{
	Use:   "delete-kx-environment",
	Short: "Deletes the kdb environment.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(finspace_deleteKxEnvironmentCmd).Standalone()

	finspace_deleteKxEnvironmentCmd.Flags().String("client-token", "", "A token that ensures idempotency.")
	finspace_deleteKxEnvironmentCmd.Flags().String("environment-id", "", "A unique identifier for the kdb environment.")
	finspace_deleteKxEnvironmentCmd.MarkFlagRequired("environment-id")
	finspaceCmd.AddCommand(finspace_deleteKxEnvironmentCmd)
}
