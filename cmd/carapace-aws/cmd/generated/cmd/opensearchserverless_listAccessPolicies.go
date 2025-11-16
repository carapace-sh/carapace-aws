package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var opensearchserverless_listAccessPoliciesCmd = &cobra.Command{
	Use:   "list-access-policies",
	Short: "Returns information about a list of OpenSearch Serverless access policies.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(opensearchserverless_listAccessPoliciesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(opensearchserverless_listAccessPoliciesCmd).Standalone()

		opensearchserverless_listAccessPoliciesCmd.Flags().String("max-results", "", "An optional parameter that specifies the maximum number of results to return.")
		opensearchserverless_listAccessPoliciesCmd.Flags().String("next-token", "", "If your initial `ListAccessPolicies` operation returns a `nextToken`, you can include the returned `nextToken` in subsequent `ListAccessPolicies` operations, which returns results in the next page.")
		opensearchserverless_listAccessPoliciesCmd.Flags().String("resource", "", "Resource filters (can be collections or indexes) that policies can apply to.")
		opensearchserverless_listAccessPoliciesCmd.Flags().String("type", "", "The type of access policy.")
		opensearchserverless_listAccessPoliciesCmd.MarkFlagRequired("type")
	})
	opensearchserverlessCmd.AddCommand(opensearchserverless_listAccessPoliciesCmd)
}
