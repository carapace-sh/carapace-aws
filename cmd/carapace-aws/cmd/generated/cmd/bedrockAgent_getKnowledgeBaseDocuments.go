package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockAgent_getKnowledgeBaseDocumentsCmd = &cobra.Command{
	Use:   "get-knowledge-base-documents",
	Short: "Retrieves specific documents from a data source that is connected to a knowledge base.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockAgent_getKnowledgeBaseDocumentsCmd).Standalone()

	bedrockAgent_getKnowledgeBaseDocumentsCmd.Flags().String("data-source-id", "", "The unique identifier of the data source that contains the documents.")
	bedrockAgent_getKnowledgeBaseDocumentsCmd.Flags().String("document-identifiers", "", "A list of objects, each of which contains information to identify a document for which to retrieve information.")
	bedrockAgent_getKnowledgeBaseDocumentsCmd.Flags().String("knowledge-base-id", "", "The unique identifier of the knowledge base that is connected to the data source.")
	bedrockAgent_getKnowledgeBaseDocumentsCmd.MarkFlagRequired("data-source-id")
	bedrockAgent_getKnowledgeBaseDocumentsCmd.MarkFlagRequired("document-identifiers")
	bedrockAgent_getKnowledgeBaseDocumentsCmd.MarkFlagRequired("knowledge-base-id")
	bedrockAgentCmd.AddCommand(bedrockAgent_getKnowledgeBaseDocumentsCmd)
}
