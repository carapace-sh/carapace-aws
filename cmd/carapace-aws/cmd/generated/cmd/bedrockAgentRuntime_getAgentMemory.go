package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockAgentRuntime_getAgentMemoryCmd = &cobra.Command{
	Use:   "get-agent-memory",
	Short: "Gets the sessions stored in the memory of the agent.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockAgentRuntime_getAgentMemoryCmd).Standalone()

	bedrockAgentRuntime_getAgentMemoryCmd.Flags().String("agent-alias-id", "", "The unique identifier of an alias of an agent.")
	bedrockAgentRuntime_getAgentMemoryCmd.Flags().String("agent-id", "", "The unique identifier of the agent to which the alias belongs.")
	bedrockAgentRuntime_getAgentMemoryCmd.Flags().String("max-items", "", "The maximum number of items to return in the response.")
	bedrockAgentRuntime_getAgentMemoryCmd.Flags().String("memory-id", "", "The unique identifier of the memory.")
	bedrockAgentRuntime_getAgentMemoryCmd.Flags().String("memory-type", "", "The type of memory.")
	bedrockAgentRuntime_getAgentMemoryCmd.Flags().String("next-token", "", "If the total number of results is greater than the maxItems value provided in the request, enter the token returned in the `nextToken` field in the response in this field to return the next batch of results.")
	bedrockAgentRuntime_getAgentMemoryCmd.MarkFlagRequired("agent-alias-id")
	bedrockAgentRuntime_getAgentMemoryCmd.MarkFlagRequired("agent-id")
	bedrockAgentRuntime_getAgentMemoryCmd.MarkFlagRequired("memory-id")
	bedrockAgentRuntime_getAgentMemoryCmd.MarkFlagRequired("memory-type")
	bedrockAgentRuntimeCmd.AddCommand(bedrockAgentRuntime_getAgentMemoryCmd)
}
