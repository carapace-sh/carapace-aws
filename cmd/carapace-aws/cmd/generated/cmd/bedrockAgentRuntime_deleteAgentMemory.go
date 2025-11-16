package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockAgentRuntime_deleteAgentMemoryCmd = &cobra.Command{
	Use:   "delete-agent-memory",
	Short: "Deletes memory from the specified memory identifier.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockAgentRuntime_deleteAgentMemoryCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(bedrockAgentRuntime_deleteAgentMemoryCmd).Standalone()

		bedrockAgentRuntime_deleteAgentMemoryCmd.Flags().String("agent-alias-id", "", "The unique identifier of an alias of an agent.")
		bedrockAgentRuntime_deleteAgentMemoryCmd.Flags().String("agent-id", "", "The unique identifier of the agent to which the alias belongs.")
		bedrockAgentRuntime_deleteAgentMemoryCmd.Flags().String("memory-id", "", "The unique identifier of the memory.")
		bedrockAgentRuntime_deleteAgentMemoryCmd.Flags().String("session-id", "", "The unique session identifier of the memory.")
		bedrockAgentRuntime_deleteAgentMemoryCmd.MarkFlagRequired("agent-alias-id")
		bedrockAgentRuntime_deleteAgentMemoryCmd.MarkFlagRequired("agent-id")
	})
	bedrockAgentRuntimeCmd.AddCommand(bedrockAgentRuntime_deleteAgentMemoryCmd)
}
