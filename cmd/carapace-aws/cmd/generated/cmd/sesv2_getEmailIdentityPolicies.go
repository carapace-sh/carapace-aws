package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sesv2_getEmailIdentityPoliciesCmd = &cobra.Command{
	Use:   "get-email-identity-policies",
	Short: "Returns the requested sending authorization policies for the given identity (an email address or a domain).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sesv2_getEmailIdentityPoliciesCmd).Standalone()

	sesv2_getEmailIdentityPoliciesCmd.Flags().String("email-identity", "", "The email identity.")
	sesv2_getEmailIdentityPoliciesCmd.MarkFlagRequired("email-identity")
	sesv2Cmd.AddCommand(sesv2_getEmailIdentityPoliciesCmd)
}
