package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sesv2_updateEmailIdentityPolicyCmd = &cobra.Command{
	Use:   "update-email-identity-policy",
	Short: "Updates the specified sending authorization policy for the given identity (an email address or a domain).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sesv2_updateEmailIdentityPolicyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sesv2_updateEmailIdentityPolicyCmd).Standalone()

		sesv2_updateEmailIdentityPolicyCmd.Flags().String("email-identity", "", "The email identity.")
		sesv2_updateEmailIdentityPolicyCmd.Flags().String("policy", "", "The text of the policy in JSON format.")
		sesv2_updateEmailIdentityPolicyCmd.Flags().String("policy-name", "", "The name of the policy.")
		sesv2_updateEmailIdentityPolicyCmd.MarkFlagRequired("email-identity")
		sesv2_updateEmailIdentityPolicyCmd.MarkFlagRequired("policy")
		sesv2_updateEmailIdentityPolicyCmd.MarkFlagRequired("policy-name")
	})
	sesv2Cmd.AddCommand(sesv2_updateEmailIdentityPolicyCmd)
}
