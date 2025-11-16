package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ses_deleteIdentityPolicyCmd = &cobra.Command{
	Use:   "delete-identity-policy",
	Short: "Deletes the specified sending authorization policy for the given identity (an email address or a domain).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ses_deleteIdentityPolicyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ses_deleteIdentityPolicyCmd).Standalone()

		ses_deleteIdentityPolicyCmd.Flags().String("identity", "", "The identity that is associated with the policy to delete.")
		ses_deleteIdentityPolicyCmd.Flags().String("policy-name", "", "The name of the policy to be deleted.")
		ses_deleteIdentityPolicyCmd.MarkFlagRequired("identity")
		ses_deleteIdentityPolicyCmd.MarkFlagRequired("policy-name")
	})
	sesCmd.AddCommand(ses_deleteIdentityPolicyCmd)
}
