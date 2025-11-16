package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockAgent_associateAgentKnowledgeBaseCmd = &cobra.Command{
	Use:   "associate-agent-knowledge-base",
	Short: "Associates a knowledge base with an agent.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockAgent_associateAgentKnowledgeBaseCmd).Standalone()

	bedrockAgent_associateAgentKnowledgeBaseCmd.Flags().String("agent-id", "", "The unique identifier of the agent with which you want to associate the knowledge base.")
	bedrockAgent_associateAgentKnowledgeBaseCmd.Flags().String("agent-version", "", "The version of the agent with which you want to associate the knowledge base.")
	bedrockAgent_associateAgentKnowledgeBaseCmd.Flags().String("description", "", "A description of what the agent should use the knowledge base for.")
	bedrockAgent_associateAgentKnowledgeBaseCmd.Flags().String("knowledge-base-id", "", "The unique identifier of the knowledge base to associate with the agent.")
	bedrockAgent_associateAgentKnowledgeBaseCmd.Flags().String("knowledge-base-state", "", "Specifies whether to use the knowledge base or not when sending an [InvokeAgent](https://docs.aws.amazon.com/bedrock/latest/APIReference/API_agent-runtime_InvokeAgent.html) request.")
	bedrockAgent_associateAgentKnowledgeBaseCmd.MarkFlagRequired("agent-id")
	bedrockAgent_associateAgentKnowledgeBaseCmd.MarkFlagRequired("agent-version")
	bedrockAgent_associateAgentKnowledgeBaseCmd.MarkFlagRequired("description")
	bedrockAgent_associateAgentKnowledgeBaseCmd.MarkFlagRequired("knowledge-base-id")
	bedrockAgentCmd.AddCommand(bedrockAgent_associateAgentKnowledgeBaseCmd)
}
