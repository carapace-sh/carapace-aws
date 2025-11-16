package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockAgentRuntime_retrieveAndGenerateCmd = &cobra.Command{
	Use:   "retrieve-and-generate",
	Short: "Queries a knowledge base and generates responses based on the retrieved results and using the specified foundation model or [inference profile](https://docs.aws.amazon.com/bedrock/latest/userguide/cross-region-inference.html).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockAgentRuntime_retrieveAndGenerateCmd).Standalone()

	bedrockAgentRuntime_retrieveAndGenerateCmd.Flags().String("input", "", "Contains the query to be made to the knowledge base.")
	bedrockAgentRuntime_retrieveAndGenerateCmd.Flags().String("retrieve-and-generate-configuration", "", "Contains configurations for the knowledge base query and retrieval process.")
	bedrockAgentRuntime_retrieveAndGenerateCmd.Flags().String("session-configuration", "", "Contains details about the session with the knowledge base.")
	bedrockAgentRuntime_retrieveAndGenerateCmd.Flags().String("session-id", "", "The unique identifier of the session.")
	bedrockAgentRuntime_retrieveAndGenerateCmd.MarkFlagRequired("input")
	bedrockAgentRuntimeCmd.AddCommand(bedrockAgentRuntime_retrieveAndGenerateCmd)
}
