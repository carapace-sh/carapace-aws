package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockAgentRuntime_optimizePromptCmd = &cobra.Command{
	Use:   "optimize-prompt",
	Short: "Optimizes a prompt for the task that you specify.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockAgentRuntime_optimizePromptCmd).Standalone()

	bedrockAgentRuntime_optimizePromptCmd.Flags().String("input", "", "Contains the prompt to optimize.")
	bedrockAgentRuntime_optimizePromptCmd.Flags().String("target-model-id", "", "The unique identifier of the model that you want to optimize the prompt for.")
	bedrockAgentRuntime_optimizePromptCmd.MarkFlagRequired("input")
	bedrockAgentRuntime_optimizePromptCmd.MarkFlagRequired("target-model-id")
	bedrockAgentRuntimeCmd.AddCommand(bedrockAgentRuntime_optimizePromptCmd)
}
