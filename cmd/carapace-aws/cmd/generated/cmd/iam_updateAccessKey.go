package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iam_updateAccessKeyCmd = &cobra.Command{
	Use:   "update-access-key",
	Short: "Changes the status of the specified access key from Active to Inactive, or vice versa.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iam_updateAccessKeyCmd).Standalone()

	iam_updateAccessKeyCmd.Flags().String("access-key-id", "", "The access key ID of the secret access key you want to update.")
	iam_updateAccessKeyCmd.Flags().String("status", "", "The status you want to assign to the secret access key.")
	iam_updateAccessKeyCmd.Flags().String("user-name", "", "The name of the user whose key you want to update.")
	iam_updateAccessKeyCmd.MarkFlagRequired("access-key-id")
	iam_updateAccessKeyCmd.MarkFlagRequired("status")
	iamCmd.AddCommand(iam_updateAccessKeyCmd)
}
