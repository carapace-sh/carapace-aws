package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sesv2_deleteEmailIdentityPolicyCmd = &cobra.Command{
	Use:   "delete-email-identity-policy",
	Short: "Deletes the specified sending authorization policy for the given identity (an email address or a domain).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sesv2_deleteEmailIdentityPolicyCmd).Standalone()

	sesv2_deleteEmailIdentityPolicyCmd.Flags().String("email-identity", "", "The email identity.")
	sesv2_deleteEmailIdentityPolicyCmd.Flags().String("policy-name", "", "The name of the policy.")
	sesv2_deleteEmailIdentityPolicyCmd.MarkFlagRequired("email-identity")
	sesv2_deleteEmailIdentityPolicyCmd.MarkFlagRequired("policy-name")
	sesv2Cmd.AddCommand(sesv2_deleteEmailIdentityPolicyCmd)
}
