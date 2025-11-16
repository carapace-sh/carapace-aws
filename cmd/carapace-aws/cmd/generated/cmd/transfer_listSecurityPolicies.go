package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var transfer_listSecurityPoliciesCmd = &cobra.Command{
	Use:   "list-security-policies",
	Short: "Lists the security policies that are attached to your servers and SFTP connectors.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(transfer_listSecurityPoliciesCmd).Standalone()

	transfer_listSecurityPoliciesCmd.Flags().String("max-results", "", "Specifies the number of security policies to return as a response to the `ListSecurityPolicies` query.")
	transfer_listSecurityPoliciesCmd.Flags().String("next-token", "", "When additional results are obtained from the `ListSecurityPolicies` command, a `NextToken` parameter is returned in the output.")
	transferCmd.AddCommand(transfer_listSecurityPoliciesCmd)
}
