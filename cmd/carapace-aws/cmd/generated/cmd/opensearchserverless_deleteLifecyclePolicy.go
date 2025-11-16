package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var opensearchserverless_deleteLifecyclePolicyCmd = &cobra.Command{
	Use:   "delete-lifecycle-policy",
	Short: "Deletes an OpenSearch Serverless lifecycle policy.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(opensearchserverless_deleteLifecyclePolicyCmd).Standalone()

	opensearchserverless_deleteLifecyclePolicyCmd.Flags().String("client-token", "", "Unique, case-sensitive identifier to ensure idempotency of the request.")
	opensearchserverless_deleteLifecyclePolicyCmd.Flags().String("name", "", "The name of the policy to delete.")
	opensearchserverless_deleteLifecyclePolicyCmd.Flags().String("type", "", "The type of lifecycle policy.")
	opensearchserverless_deleteLifecyclePolicyCmd.MarkFlagRequired("name")
	opensearchserverless_deleteLifecyclePolicyCmd.MarkFlagRequired("type")
	opensearchserverlessCmd.AddCommand(opensearchserverless_deleteLifecyclePolicyCmd)
}
