package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockAgentRuntime_rerankCmd = &cobra.Command{
	Use:   "rerank",
	Short: "Reranks the relevance of sources based on queries.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockAgentRuntime_rerankCmd).Standalone()

	bedrockAgentRuntime_rerankCmd.Flags().String("next-token", "", "If the total number of results was greater than could fit in a response, a token is returned in the `nextToken` field.")
	bedrockAgentRuntime_rerankCmd.Flags().String("queries", "", "An array of objects, each of which contains information about a query to submit to the reranker model.")
	bedrockAgentRuntime_rerankCmd.Flags().String("reranking-configuration", "", "Contains configurations for reranking.")
	bedrockAgentRuntime_rerankCmd.Flags().String("sources", "", "An array of objects, each of which contains information about the sources to rerank.")
	bedrockAgentRuntime_rerankCmd.MarkFlagRequired("queries")
	bedrockAgentRuntime_rerankCmd.MarkFlagRequired("reranking-configuration")
	bedrockAgentRuntime_rerankCmd.MarkFlagRequired("sources")
	bedrockAgentRuntimeCmd.AddCommand(bedrockAgentRuntime_rerankCmd)
}
