package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sesv2_createEmailIdentityPolicyCmd = &cobra.Command{
	Use:   "create-email-identity-policy",
	Short: "Creates the specified sending authorization policy for the given identity (an email address or a domain).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sesv2_createEmailIdentityPolicyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sesv2_createEmailIdentityPolicyCmd).Standalone()

		sesv2_createEmailIdentityPolicyCmd.Flags().String("email-identity", "", "The email identity.")
		sesv2_createEmailIdentityPolicyCmd.Flags().String("policy", "", "The text of the policy in JSON format.")
		sesv2_createEmailIdentityPolicyCmd.Flags().String("policy-name", "", "The name of the policy.")
		sesv2_createEmailIdentityPolicyCmd.MarkFlagRequired("email-identity")
		sesv2_createEmailIdentityPolicyCmd.MarkFlagRequired("policy")
		sesv2_createEmailIdentityPolicyCmd.MarkFlagRequired("policy-name")
	})
	sesv2Cmd.AddCommand(sesv2_createEmailIdentityPolicyCmd)
}
