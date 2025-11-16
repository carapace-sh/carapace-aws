package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockAgentRuntime_retrieveCmd = &cobra.Command{
	Use:   "retrieve",
	Short: "Queries a knowledge base and retrieves information from it.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockAgentRuntime_retrieveCmd).Standalone()

	bedrockAgentRuntime_retrieveCmd.Flags().String("guardrail-configuration", "", "Guardrail settings.")
	bedrockAgentRuntime_retrieveCmd.Flags().String("knowledge-base-id", "", "The unique identifier of the knowledge base to query.")
	bedrockAgentRuntime_retrieveCmd.Flags().String("next-token", "", "If there are more results than can fit in the response, the response returns a `nextToken`.")
	bedrockAgentRuntime_retrieveCmd.Flags().String("retrieval-configuration", "", "Contains configurations for the knowledge base query and retrieval process.")
	bedrockAgentRuntime_retrieveCmd.Flags().String("retrieval-query", "", "Contains the query to send the knowledge base.")
	bedrockAgentRuntime_retrieveCmd.MarkFlagRequired("knowledge-base-id")
	bedrockAgentRuntime_retrieveCmd.MarkFlagRequired("retrieval-query")
	bedrockAgentRuntimeCmd.AddCommand(bedrockAgentRuntime_retrieveCmd)
}
