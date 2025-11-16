package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockAgent_getAgentKnowledgeBaseCmd = &cobra.Command{
	Use:   "get-agent-knowledge-base",
	Short: "Gets information about a knowledge base associated with an agent.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockAgent_getAgentKnowledgeBaseCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(bedrockAgent_getAgentKnowledgeBaseCmd).Standalone()

		bedrockAgent_getAgentKnowledgeBaseCmd.Flags().String("agent-id", "", "The unique identifier of the agent with which the knowledge base is associated.")
		bedrockAgent_getAgentKnowledgeBaseCmd.Flags().String("agent-version", "", "The version of the agent with which the knowledge base is associated.")
		bedrockAgent_getAgentKnowledgeBaseCmd.Flags().String("knowledge-base-id", "", "The unique identifier of the knowledge base associated with the agent.")
		bedrockAgent_getAgentKnowledgeBaseCmd.MarkFlagRequired("agent-id")
		bedrockAgent_getAgentKnowledgeBaseCmd.MarkFlagRequired("agent-version")
		bedrockAgent_getAgentKnowledgeBaseCmd.MarkFlagRequired("knowledge-base-id")
	})
	bedrockAgentCmd.AddCommand(bedrockAgent_getAgentKnowledgeBaseCmd)
}
