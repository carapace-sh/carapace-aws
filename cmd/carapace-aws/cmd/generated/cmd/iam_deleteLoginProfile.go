package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iam_deleteLoginProfileCmd = &cobra.Command{
	Use:   "delete-login-profile",
	Short: "Deletes the password for the specified IAM user or root user, For more information, see [Managing passwords for IAM users](https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_passwords_admin-change-user.html).\n\nYou can use the CLI, the Amazon Web Services API, or the **Users** page in the IAM console to delete a password for any IAM user.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iam_deleteLoginProfileCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iam_deleteLoginProfileCmd).Standalone()

		iam_deleteLoginProfileCmd.Flags().String("user-name", "", "The name of the user whose password you want to delete.")
	})
	iamCmd.AddCommand(iam_deleteLoginProfileCmd)
}
