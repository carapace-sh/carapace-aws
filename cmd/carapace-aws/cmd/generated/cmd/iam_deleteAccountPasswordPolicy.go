package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iam_deleteAccountPasswordPolicyCmd = &cobra.Command{
	Use:   "delete-account-password-policy",
	Short: "Deletes the password policy for the Amazon Web Services account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iam_deleteAccountPasswordPolicyCmd).Standalone()

	iamCmd.AddCommand(iam_deleteAccountPasswordPolicyCmd)
}
