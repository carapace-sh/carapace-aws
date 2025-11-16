package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var opensearchserverless_updateLifecyclePolicyCmd = &cobra.Command{
	Use:   "update-lifecycle-policy",
	Short: "Updates an OpenSearch Serverless access policy.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(opensearchserverless_updateLifecyclePolicyCmd).Standalone()

	opensearchserverless_updateLifecyclePolicyCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier to ensure idempotency of the request.")
	opensearchserverless_updateLifecyclePolicyCmd.Flags().String("description", "", "A description of the lifecycle policy.")
	opensearchserverless_updateLifecyclePolicyCmd.Flags().String("name", "", "The name of the policy.")
	opensearchserverless_updateLifecyclePolicyCmd.Flags().String("policy", "", "The JSON policy document to use as the content for the lifecycle policy.")
	opensearchserverless_updateLifecyclePolicyCmd.Flags().String("policy-version", "", "The version of the policy being updated.")
	opensearchserverless_updateLifecyclePolicyCmd.Flags().String("type", "", "The type of lifecycle policy.")
	opensearchserverless_updateLifecyclePolicyCmd.MarkFlagRequired("name")
	opensearchserverless_updateLifecyclePolicyCmd.MarkFlagRequired("policy-version")
	opensearchserverless_updateLifecyclePolicyCmd.MarkFlagRequired("type")
	opensearchserverlessCmd.AddCommand(opensearchserverless_updateLifecyclePolicyCmd)
}
