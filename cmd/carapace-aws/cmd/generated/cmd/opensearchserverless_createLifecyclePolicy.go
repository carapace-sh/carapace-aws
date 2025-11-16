package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var opensearchserverless_createLifecyclePolicyCmd = &cobra.Command{
	Use:   "create-lifecycle-policy",
	Short: "Creates a lifecyle policy to be applied to OpenSearch Serverless indexes.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(opensearchserverless_createLifecyclePolicyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(opensearchserverless_createLifecyclePolicyCmd).Standalone()

		opensearchserverless_createLifecyclePolicyCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier to ensure idempotency of the request.")
		opensearchserverless_createLifecyclePolicyCmd.Flags().String("description", "", "A description of the lifecycle policy.")
		opensearchserverless_createLifecyclePolicyCmd.Flags().String("name", "", "The name of the lifecycle policy.")
		opensearchserverless_createLifecyclePolicyCmd.Flags().String("policy", "", "The JSON policy document to use as the content for the lifecycle policy.")
		opensearchserverless_createLifecyclePolicyCmd.Flags().String("type", "", "The type of lifecycle policy.")
		opensearchserverless_createLifecyclePolicyCmd.MarkFlagRequired("name")
		opensearchserverless_createLifecyclePolicyCmd.MarkFlagRequired("policy")
		opensearchserverless_createLifecyclePolicyCmd.MarkFlagRequired("type")
	})
	opensearchserverlessCmd.AddCommand(opensearchserverless_createLifecyclePolicyCmd)
}
