package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var opensearchserverless_batchGetLifecyclePolicyCmd = &cobra.Command{
	Use:   "batch-get-lifecycle-policy",
	Short: "Returns one or more configured OpenSearch Serverless lifecycle policies.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(opensearchserverless_batchGetLifecyclePolicyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(opensearchserverless_batchGetLifecyclePolicyCmd).Standalone()

		opensearchserverless_batchGetLifecyclePolicyCmd.Flags().String("identifiers", "", "The unique identifiers of policy types and policy names.")
		opensearchserverless_batchGetLifecyclePolicyCmd.MarkFlagRequired("identifiers")
	})
	opensearchserverlessCmd.AddCommand(opensearchserverless_batchGetLifecyclePolicyCmd)
}
