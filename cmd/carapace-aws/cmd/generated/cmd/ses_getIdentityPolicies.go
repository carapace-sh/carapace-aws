package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ses_getIdentityPoliciesCmd = &cobra.Command{
	Use:   "get-identity-policies",
	Short: "Returns the requested sending authorization policies for the given identity (an email address or a domain).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ses_getIdentityPoliciesCmd).Standalone()

	ses_getIdentityPoliciesCmd.Flags().String("identity", "", "The identity for which the policies are retrieved.")
	ses_getIdentityPoliciesCmd.Flags().String("policy-names", "", "A list of the names of policies to be retrieved.")
	ses_getIdentityPoliciesCmd.MarkFlagRequired("identity")
	ses_getIdentityPoliciesCmd.MarkFlagRequired("policy-names")
	sesCmd.AddCommand(ses_getIdentityPoliciesCmd)
}
