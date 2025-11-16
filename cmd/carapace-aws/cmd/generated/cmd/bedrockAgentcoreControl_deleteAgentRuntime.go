package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockAgentcoreControl_deleteAgentRuntimeCmd = &cobra.Command{
	Use:   "delete-agent-runtime",
	Short: "Deletes an Amazon Bedrock AgentCore Runtime.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockAgentcoreControl_deleteAgentRuntimeCmd).Standalone()

	bedrockAgentcoreControl_deleteAgentRuntimeCmd.Flags().String("agent-runtime-id", "", "The unique identifier of the AgentCore Runtime to delete.")
	bedrockAgentcoreControl_deleteAgentRuntimeCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier to ensure that the operation completes no more than one time.")
	bedrockAgentcoreControl_deleteAgentRuntimeCmd.MarkFlagRequired("agent-runtime-id")
	bedrockAgentcoreControlCmd.AddCommand(bedrockAgentcoreControl_deleteAgentRuntimeCmd)
}
