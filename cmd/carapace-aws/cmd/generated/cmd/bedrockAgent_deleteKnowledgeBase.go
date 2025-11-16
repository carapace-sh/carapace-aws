package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockAgent_deleteKnowledgeBaseCmd = &cobra.Command{
	Use:   "delete-knowledge-base",
	Short: "Deletes a knowledge base.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockAgent_deleteKnowledgeBaseCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(bedrockAgent_deleteKnowledgeBaseCmd).Standalone()

		bedrockAgent_deleteKnowledgeBaseCmd.Flags().String("knowledge-base-id", "", "The unique identifier of the knowledge base to delete.")
		bedrockAgent_deleteKnowledgeBaseCmd.MarkFlagRequired("knowledge-base-id")
	})
	bedrockAgentCmd.AddCommand(bedrockAgent_deleteKnowledgeBaseCmd)
}
