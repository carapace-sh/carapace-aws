package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iam_createLoginProfileCmd = &cobra.Command{
	Use:   "create-login-profile",
	Short: "Creates a password for the specified IAM user.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iam_createLoginProfileCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iam_createLoginProfileCmd).Standalone()

		iam_createLoginProfileCmd.Flags().String("password", "", "The new password for the user.")
		iam_createLoginProfileCmd.Flags().String("password-reset-required", "", "Specifies whether the user is required to set a new password on next sign-in.")
		iam_createLoginProfileCmd.Flags().String("user-name", "", "The name of the IAM user to create a password for.")
	})
	iamCmd.AddCommand(iam_createLoginProfileCmd)
}
