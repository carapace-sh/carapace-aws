package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ses_putIdentityPolicyCmd = &cobra.Command{
	Use:   "put-identity-policy",
	Short: "Adds or updates a sending authorization policy for the specified identity (an email address or a domain).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ses_putIdentityPolicyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ses_putIdentityPolicyCmd).Standalone()

		ses_putIdentityPolicyCmd.Flags().String("identity", "", "The identity to which that the policy applies.")
		ses_putIdentityPolicyCmd.Flags().String("policy", "", "The text of the policy in JSON format.")
		ses_putIdentityPolicyCmd.Flags().String("policy-name", "", "The name of the policy.")
		ses_putIdentityPolicyCmd.MarkFlagRequired("identity")
		ses_putIdentityPolicyCmd.MarkFlagRequired("policy")
		ses_putIdentityPolicyCmd.MarkFlagRequired("policy-name")
	})
	sesCmd.AddCommand(ses_putIdentityPolicyCmd)
}
