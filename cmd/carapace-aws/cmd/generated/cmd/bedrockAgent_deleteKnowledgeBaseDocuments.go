package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockAgent_deleteKnowledgeBaseDocumentsCmd = &cobra.Command{
	Use:   "delete-knowledge-base-documents",
	Short: "Deletes documents from a data source and syncs the changes to the knowledge base that is connected to it.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockAgent_deleteKnowledgeBaseDocumentsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(bedrockAgent_deleteKnowledgeBaseDocumentsCmd).Standalone()

		bedrockAgent_deleteKnowledgeBaseDocumentsCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier to ensure that the API request completes no more than one time.")
		bedrockAgent_deleteKnowledgeBaseDocumentsCmd.Flags().String("data-source-id", "", "The unique identifier of the data source that contains the documents.")
		bedrockAgent_deleteKnowledgeBaseDocumentsCmd.Flags().String("document-identifiers", "", "A list of objects, each of which contains information to identify a document to delete.")
		bedrockAgent_deleteKnowledgeBaseDocumentsCmd.Flags().String("knowledge-base-id", "", "The unique identifier of the knowledge base that is connected to the data source.")
		bedrockAgent_deleteKnowledgeBaseDocumentsCmd.MarkFlagRequired("data-source-id")
		bedrockAgent_deleteKnowledgeBaseDocumentsCmd.MarkFlagRequired("document-identifiers")
		bedrockAgent_deleteKnowledgeBaseDocumentsCmd.MarkFlagRequired("knowledge-base-id")
	})
	bedrockAgentCmd.AddCommand(bedrockAgent_deleteKnowledgeBaseDocumentsCmd)
}
