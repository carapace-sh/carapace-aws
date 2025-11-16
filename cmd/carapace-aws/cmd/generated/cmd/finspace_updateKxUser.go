package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var finspace_updateKxUserCmd = &cobra.Command{
	Use:   "update-kx-user",
	Short: "Updates the user details.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(finspace_updateKxUserCmd).Standalone()

	finspace_updateKxUserCmd.Flags().String("client-token", "", "A token that ensures idempotency.")
	finspace_updateKxUserCmd.Flags().String("environment-id", "", "A unique identifier for the kdb environment.")
	finspace_updateKxUserCmd.Flags().String("iam-role", "", "The IAM role ARN that is associated with the user.")
	finspace_updateKxUserCmd.Flags().String("user-name", "", "A unique identifier for the user.")
	finspace_updateKxUserCmd.MarkFlagRequired("environment-id")
	finspace_updateKxUserCmd.MarkFlagRequired("iam-role")
	finspace_updateKxUserCmd.MarkFlagRequired("user-name")
	finspaceCmd.AddCommand(finspace_updateKxUserCmd)
}
