package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var finspace_createKxUserCmd = &cobra.Command{
	Use:   "create-kx-user",
	Short: "Creates a user in FinSpace kdb environment with an associated IAM role.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(finspace_createKxUserCmd).Standalone()

	finspace_createKxUserCmd.Flags().String("client-token", "", "A token that ensures idempotency.")
	finspace_createKxUserCmd.Flags().String("environment-id", "", "A unique identifier for the kdb environment where you want to create a user.")
	finspace_createKxUserCmd.Flags().String("iam-role", "", "The IAM role ARN that will be associated with the user.")
	finspace_createKxUserCmd.Flags().String("tags", "", "A list of key-value pairs to label the user.")
	finspace_createKxUserCmd.Flags().String("user-name", "", "A unique identifier for the user.")
	finspace_createKxUserCmd.MarkFlagRequired("environment-id")
	finspace_createKxUserCmd.MarkFlagRequired("iam-role")
	finspace_createKxUserCmd.MarkFlagRequired("user-name")
	finspaceCmd.AddCommand(finspace_createKxUserCmd)
}
