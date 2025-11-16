package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockAgent_getKnowledgeBaseCmd = &cobra.Command{
	Use:   "get-knowledge-base",
	Short: "Gets information about a knoweldge base.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockAgent_getKnowledgeBaseCmd).Standalone()

	bedrockAgent_getKnowledgeBaseCmd.Flags().String("knowledge-base-id", "", "The unique identifier of the knowledge base you want to get information on.")
	bedrockAgent_getKnowledgeBaseCmd.MarkFlagRequired("knowledge-base-id")
	bedrockAgentCmd.AddCommand(bedrockAgent_getKnowledgeBaseCmd)
}
