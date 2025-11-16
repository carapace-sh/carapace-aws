package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var opensearchserverless_listSecurityPoliciesCmd = &cobra.Command{
	Use:   "list-security-policies",
	Short: "Returns information about configured OpenSearch Serverless security policies.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(opensearchserverless_listSecurityPoliciesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(opensearchserverless_listSecurityPoliciesCmd).Standalone()

		opensearchserverless_listSecurityPoliciesCmd.Flags().String("max-results", "", "An optional parameter that specifies the maximum number of results to return.")
		opensearchserverless_listSecurityPoliciesCmd.Flags().String("next-token", "", "If your initial `ListSecurityPolicies` operation returns a `nextToken`, you can include the returned `nextToken` in subsequent `ListSecurityPolicies` operations, which returns results in the next page.")
		opensearchserverless_listSecurityPoliciesCmd.Flags().String("resource", "", "Resource filters (can be collection or indexes) that policies can apply to.")
		opensearchserverless_listSecurityPoliciesCmd.Flags().String("type", "", "The type of policy.")
		opensearchserverless_listSecurityPoliciesCmd.MarkFlagRequired("type")
	})
	opensearchserverlessCmd.AddCommand(opensearchserverless_listSecurityPoliciesCmd)
}
