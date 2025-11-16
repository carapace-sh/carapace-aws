package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var opensearchserverless_createAccessPolicyCmd = &cobra.Command{
	Use:   "create-access-policy",
	Short: "Creates a data access policy for OpenSearch Serverless.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(opensearchserverless_createAccessPolicyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(opensearchserverless_createAccessPolicyCmd).Standalone()

		opensearchserverless_createAccessPolicyCmd.Flags().String("client-token", "", "Unique, case-sensitive identifier to ensure idempotency of the request.")
		opensearchserverless_createAccessPolicyCmd.Flags().String("description", "", "A description of the policy.")
		opensearchserverless_createAccessPolicyCmd.Flags().String("name", "", "The name of the policy.")
		opensearchserverless_createAccessPolicyCmd.Flags().String("policy", "", "The JSON policy document to use as the content for the policy.")
		opensearchserverless_createAccessPolicyCmd.Flags().String("type", "", "The type of policy.")
		opensearchserverless_createAccessPolicyCmd.MarkFlagRequired("name")
		opensearchserverless_createAccessPolicyCmd.MarkFlagRequired("policy")
		opensearchserverless_createAccessPolicyCmd.MarkFlagRequired("type")
	})
	opensearchserverlessCmd.AddCommand(opensearchserverless_createAccessPolicyCmd)
}
