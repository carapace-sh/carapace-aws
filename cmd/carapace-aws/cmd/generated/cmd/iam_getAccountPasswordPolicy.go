package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iam_getAccountPasswordPolicyCmd = &cobra.Command{
	Use:   "get-account-password-policy",
	Short: "Retrieves the password policy for the Amazon Web Services account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iam_getAccountPasswordPolicyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iam_getAccountPasswordPolicyCmd).Standalone()

	})
	iamCmd.AddCommand(iam_getAccountPasswordPolicyCmd)
}
