package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockAgentcoreControl_deleteMemoryCmd = &cobra.Command{
	Use:   "delete-memory",
	Short: "Deletes an Amazon Bedrock AgentCore Memory resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockAgentcoreControl_deleteMemoryCmd).Standalone()

	bedrockAgentcoreControl_deleteMemoryCmd.Flags().String("client-token", "", "A client token is used for keeping track of idempotent requests.")
	bedrockAgentcoreControl_deleteMemoryCmd.Flags().String("memory-id", "", "The unique identifier of the memory to delete.")
	bedrockAgentcoreControl_deleteMemoryCmd.MarkFlagRequired("memory-id")
	bedrockAgentcoreControlCmd.AddCommand(bedrockAgentcoreControl_deleteMemoryCmd)
}
