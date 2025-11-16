package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockAgentRuntime_retrieveAndGenerateStreamCmd = &cobra.Command{
	Use:   "retrieve-and-generate-stream",
	Short: "Queries a knowledge base and generates responses based on the retrieved results, with output in streaming format.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockAgentRuntime_retrieveAndGenerateStreamCmd).Standalone()

	bedrockAgentRuntime_retrieveAndGenerateStreamCmd.Flags().String("input", "", "Contains the query to be made to the knowledge base.")
	bedrockAgentRuntime_retrieveAndGenerateStreamCmd.Flags().String("retrieve-and-generate-configuration", "", "Contains configurations for the knowledge base query and retrieval process.")
	bedrockAgentRuntime_retrieveAndGenerateStreamCmd.Flags().String("session-configuration", "", "Contains details about the session with the knowledge base.")
	bedrockAgentRuntime_retrieveAndGenerateStreamCmd.Flags().String("session-id", "", "The unique identifier of the session.")
	bedrockAgentRuntime_retrieveAndGenerateStreamCmd.MarkFlagRequired("input")
	bedrockAgentRuntimeCmd.AddCommand(bedrockAgentRuntime_retrieveAndGenerateStreamCmd)
}
