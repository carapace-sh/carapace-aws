package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iam_updateAccountPasswordPolicyCmd = &cobra.Command{
	Use:   "update-account-password-policy",
	Short: "Updates the password policy settings for the Amazon Web Services account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iam_updateAccountPasswordPolicyCmd).Standalone()

	iam_updateAccountPasswordPolicyCmd.Flags().String("allow-users-to-change-password", "", "Allows all IAM users in your account to use the Amazon Web Services Management Console to change their own passwords.")
	iam_updateAccountPasswordPolicyCmd.Flags().String("hard-expiry", "", "Prevents IAM users who are accessing the account via the Amazon Web Services Management Console from setting a new console password after their password has expired.")
	iam_updateAccountPasswordPolicyCmd.Flags().String("max-password-age", "", "The number of days that an IAM user password is valid.")
	iam_updateAccountPasswordPolicyCmd.Flags().String("minimum-password-length", "", "The minimum number of characters allowed in an IAM user password.")
	iam_updateAccountPasswordPolicyCmd.Flags().String("password-reuse-prevention", "", "Specifies the number of previous passwords that IAM users are prevented from reusing.")
	iam_updateAccountPasswordPolicyCmd.Flags().String("require-lowercase-characters", "", "Specifies whether IAM user passwords must contain at least one lowercase character from the ISO basic Latin alphabet (a to z).")
	iam_updateAccountPasswordPolicyCmd.Flags().String("require-numbers", "", "Specifies whether IAM user passwords must contain at least one numeric character (0 to 9).")
	iam_updateAccountPasswordPolicyCmd.Flags().String("require-symbols", "", "Specifies whether IAM user passwords must contain at least one of the following non-alphanumeric characters:")
	iam_updateAccountPasswordPolicyCmd.Flags().String("require-uppercase-characters", "", "Specifies whether IAM user passwords must contain at least one uppercase character from the ISO basic Latin alphabet (A to Z).")
	iamCmd.AddCommand(iam_updateAccountPasswordPolicyCmd)
}
