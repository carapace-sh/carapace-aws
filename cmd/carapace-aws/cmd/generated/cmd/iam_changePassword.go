package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iam_changePasswordCmd = &cobra.Command{
	Use:   "change-password",
	Short: "Changes the password of the IAM user who is calling this operation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iam_changePasswordCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iam_changePasswordCmd).Standalone()

		iam_changePasswordCmd.Flags().String("new-password", "", "The new password.")
		iam_changePasswordCmd.Flags().String("old-password", "", "The IAM user's current password.")
		iam_changePasswordCmd.MarkFlagRequired("new-password")
		iam_changePasswordCmd.MarkFlagRequired("old-password")
	})
	iamCmd.AddCommand(iam_changePasswordCmd)
}
