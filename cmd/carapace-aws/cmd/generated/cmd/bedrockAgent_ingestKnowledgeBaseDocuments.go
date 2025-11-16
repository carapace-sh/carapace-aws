package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockAgent_ingestKnowledgeBaseDocumentsCmd = &cobra.Command{
	Use:   "ingest-knowledge-base-documents",
	Short: "Ingests documents directly into the knowledge base that is connected to the data source.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockAgent_ingestKnowledgeBaseDocumentsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(bedrockAgent_ingestKnowledgeBaseDocumentsCmd).Standalone()

		bedrockAgent_ingestKnowledgeBaseDocumentsCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier to ensure that the API request completes no more than one time.")
		bedrockAgent_ingestKnowledgeBaseDocumentsCmd.Flags().String("data-source-id", "", "The unique identifier of the data source connected to the knowledge base that you're adding documents to.")
		bedrockAgent_ingestKnowledgeBaseDocumentsCmd.Flags().String("documents", "", "A list of objects, each of which contains information about the documents to add.")
		bedrockAgent_ingestKnowledgeBaseDocumentsCmd.Flags().String("knowledge-base-id", "", "The unique identifier of the knowledge base to ingest the documents into.")
		bedrockAgent_ingestKnowledgeBaseDocumentsCmd.MarkFlagRequired("data-source-id")
		bedrockAgent_ingestKnowledgeBaseDocumentsCmd.MarkFlagRequired("documents")
		bedrockAgent_ingestKnowledgeBaseDocumentsCmd.MarkFlagRequired("knowledge-base-id")
	})
	bedrockAgentCmd.AddCommand(bedrockAgent_ingestKnowledgeBaseDocumentsCmd)
}
