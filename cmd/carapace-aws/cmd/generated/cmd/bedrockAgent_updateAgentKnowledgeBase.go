package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockAgent_updateAgentKnowledgeBaseCmd = &cobra.Command{
	Use:   "update-agent-knowledge-base",
	Short: "Updates the configuration for a knowledge base that has been associated with an agent.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockAgent_updateAgentKnowledgeBaseCmd).Standalone()

	bedrockAgent_updateAgentKnowledgeBaseCmd.Flags().String("agent-id", "", "The unique identifier of the agent associated with the knowledge base that you want to update.")
	bedrockAgent_updateAgentKnowledgeBaseCmd.Flags().String("agent-version", "", "The version of the agent associated with the knowledge base that you want to update.")
	bedrockAgent_updateAgentKnowledgeBaseCmd.Flags().String("description", "", "Specifies a new description for the knowledge base associated with an agent.")
	bedrockAgent_updateAgentKnowledgeBaseCmd.Flags().String("knowledge-base-id", "", "The unique identifier of the knowledge base that has been associated with an agent.")
	bedrockAgent_updateAgentKnowledgeBaseCmd.Flags().String("knowledge-base-state", "", "Specifies whether the agent uses the knowledge base or not when sending an [InvokeAgent](https://docs.aws.amazon.com/bedrock/latest/APIReference/API_agent-runtime_InvokeAgent.html) request.")
	bedrockAgent_updateAgentKnowledgeBaseCmd.MarkFlagRequired("agent-id")
	bedrockAgent_updateAgentKnowledgeBaseCmd.MarkFlagRequired("agent-version")
	bedrockAgent_updateAgentKnowledgeBaseCmd.MarkFlagRequired("knowledge-base-id")
	bedrockAgentCmd.AddCommand(bedrockAgent_updateAgentKnowledgeBaseCmd)
}
