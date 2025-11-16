package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ses_listIdentityPoliciesCmd = &cobra.Command{
	Use:   "list-identity-policies",
	Short: "Returns a list of sending authorization policies that are attached to the given identity (an email address or a domain).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ses_listIdentityPoliciesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ses_listIdentityPoliciesCmd).Standalone()

		ses_listIdentityPoliciesCmd.Flags().String("identity", "", "The identity that is associated with the policy for which the policies are listed.")
		ses_listIdentityPoliciesCmd.MarkFlagRequired("identity")
	})
	sesCmd.AddCommand(ses_listIdentityPoliciesCmd)
}
