package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var opensearchserverless_batchGetEffectiveLifecyclePolicyCmd = &cobra.Command{
	Use:   "batch-get-effective-lifecycle-policy",
	Short: "Returns a list of successful and failed retrievals for the OpenSearch Serverless indexes.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(opensearchserverless_batchGetEffectiveLifecyclePolicyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(opensearchserverless_batchGetEffectiveLifecyclePolicyCmd).Standalone()

		opensearchserverless_batchGetEffectiveLifecyclePolicyCmd.Flags().String("resource-identifiers", "", "The unique identifiers of policy types and resource names.")
		opensearchserverless_batchGetEffectiveLifecyclePolicyCmd.MarkFlagRequired("resource-identifiers")
	})
	opensearchserverlessCmd.AddCommand(opensearchserverless_batchGetEffectiveLifecyclePolicyCmd)
}
