package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockAgent_disassociateAgentKnowledgeBaseCmd = &cobra.Command{
	Use:   "disassociate-agent-knowledge-base",
	Short: "Disassociates a knowledge base from an agent.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockAgent_disassociateAgentKnowledgeBaseCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(bedrockAgent_disassociateAgentKnowledgeBaseCmd).Standalone()

		bedrockAgent_disassociateAgentKnowledgeBaseCmd.Flags().String("agent-id", "", "The unique identifier of the agent from which to disassociate the knowledge base.")
		bedrockAgent_disassociateAgentKnowledgeBaseCmd.Flags().String("agent-version", "", "The version of the agent from which to disassociate the knowledge base.")
		bedrockAgent_disassociateAgentKnowledgeBaseCmd.Flags().String("knowledge-base-id", "", "The unique identifier of the knowledge base to disassociate.")
		bedrockAgent_disassociateAgentKnowledgeBaseCmd.MarkFlagRequired("agent-id")
		bedrockAgent_disassociateAgentKnowledgeBaseCmd.MarkFlagRequired("agent-version")
		bedrockAgent_disassociateAgentKnowledgeBaseCmd.MarkFlagRequired("knowledge-base-id")
	})
	bedrockAgentCmd.AddCommand(bedrockAgent_disassociateAgentKnowledgeBaseCmd)
}
