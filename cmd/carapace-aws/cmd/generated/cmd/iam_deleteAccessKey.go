package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iam_deleteAccessKeyCmd = &cobra.Command{
	Use:   "delete-access-key",
	Short: "Deletes the access key pair associated with the specified IAM user.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iam_deleteAccessKeyCmd).Standalone()

	iam_deleteAccessKeyCmd.Flags().String("access-key-id", "", "The access key ID for the access key ID and secret access key you want to delete.")
	iam_deleteAccessKeyCmd.Flags().String("user-name", "", "The name of the user whose access key pair you want to delete.")
	iam_deleteAccessKeyCmd.MarkFlagRequired("access-key-id")
	iamCmd.AddCommand(iam_deleteAccessKeyCmd)
}
