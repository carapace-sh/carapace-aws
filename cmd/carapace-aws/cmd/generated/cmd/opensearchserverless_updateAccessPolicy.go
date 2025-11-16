package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var opensearchserverless_updateAccessPolicyCmd = &cobra.Command{
	Use:   "update-access-policy",
	Short: "Updates an OpenSearch Serverless access policy.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(opensearchserverless_updateAccessPolicyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(opensearchserverless_updateAccessPolicyCmd).Standalone()

		opensearchserverless_updateAccessPolicyCmd.Flags().String("client-token", "", "Unique, case-sensitive identifier to ensure idempotency of the request.")
		opensearchserverless_updateAccessPolicyCmd.Flags().String("description", "", "A description of the policy.")
		opensearchserverless_updateAccessPolicyCmd.Flags().String("name", "", "The name of the policy.")
		opensearchserverless_updateAccessPolicyCmd.Flags().String("policy", "", "The JSON policy document to use as the content for the policy.")
		opensearchserverless_updateAccessPolicyCmd.Flags().String("policy-version", "", "The version of the policy being updated.")
		opensearchserverless_updateAccessPolicyCmd.Flags().String("type", "", "The type of policy.")
		opensearchserverless_updateAccessPolicyCmd.MarkFlagRequired("name")
		opensearchserverless_updateAccessPolicyCmd.MarkFlagRequired("policy-version")
		opensearchserverless_updateAccessPolicyCmd.MarkFlagRequired("type")
	})
	opensearchserverlessCmd.AddCommand(opensearchserverless_updateAccessPolicyCmd)
}
