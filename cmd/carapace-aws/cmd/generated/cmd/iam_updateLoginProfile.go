package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iam_updateLoginProfileCmd = &cobra.Command{
	Use:   "update-login-profile",
	Short: "Changes the password for the specified IAM user.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iam_updateLoginProfileCmd).Standalone()

	iam_updateLoginProfileCmd.Flags().String("password", "", "The new password for the specified IAM user.")
	iam_updateLoginProfileCmd.Flags().String("password-reset-required", "", "Allows this new password to be used only once by requiring the specified IAM user to set a new password on next sign-in.")
	iam_updateLoginProfileCmd.Flags().String("user-name", "", "The name of the user whose password you want to update.")
	iam_updateLoginProfileCmd.MarkFlagRequired("user-name")
	iamCmd.AddCommand(iam_updateLoginProfileCmd)
}
