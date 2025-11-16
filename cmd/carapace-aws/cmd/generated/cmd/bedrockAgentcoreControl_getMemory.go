package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockAgentcoreControl_getMemoryCmd = &cobra.Command{
	Use:   "get-memory",
	Short: "Retrieve an existing Amazon Bedrock AgentCore Memory resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockAgentcoreControl_getMemoryCmd).Standalone()

	bedrockAgentcoreControl_getMemoryCmd.Flags().String("memory-id", "", "The unique identifier of the memory to retrieve.")
	bedrockAgentcoreControl_getMemoryCmd.MarkFlagRequired("memory-id")
	bedrockAgentcoreControlCmd.AddCommand(bedrockAgentcoreControl_getMemoryCmd)
}
