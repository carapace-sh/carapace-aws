package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var opensearchserverless_listLifecyclePoliciesCmd = &cobra.Command{
	Use:   "list-lifecycle-policies",
	Short: "Returns a list of OpenSearch Serverless lifecycle policies.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(opensearchserverless_listLifecyclePoliciesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(opensearchserverless_listLifecyclePoliciesCmd).Standalone()

		opensearchserverless_listLifecyclePoliciesCmd.Flags().String("max-results", "", "An optional parameter that specifies the maximum number of results to return.")
		opensearchserverless_listLifecyclePoliciesCmd.Flags().String("next-token", "", "If your initial `ListLifecyclePolicies` operation returns a `nextToken`, you can include the returned `nextToken` in subsequent `ListLifecyclePolicies` operations, which returns results in the next page.")
		opensearchserverless_listLifecyclePoliciesCmd.Flags().String("resources", "", "Resource filters that policies can apply to.")
		opensearchserverless_listLifecyclePoliciesCmd.Flags().String("type", "", "The type of lifecycle policy.")
		opensearchserverless_listLifecyclePoliciesCmd.MarkFlagRequired("type")
	})
	opensearchserverlessCmd.AddCommand(opensearchserverless_listLifecyclePoliciesCmd)
}
