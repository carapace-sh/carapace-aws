package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockAgent_listKnowledgeBaseDocumentsCmd = &cobra.Command{
	Use:   "list-knowledge-base-documents",
	Short: "Retrieves all the documents contained in a data source that is connected to a knowledge base.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockAgent_listKnowledgeBaseDocumentsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(bedrockAgent_listKnowledgeBaseDocumentsCmd).Standalone()

		bedrockAgent_listKnowledgeBaseDocumentsCmd.Flags().String("data-source-id", "", "The unique identifier of the data source that contains the documents.")
		bedrockAgent_listKnowledgeBaseDocumentsCmd.Flags().String("knowledge-base-id", "", "The unique identifier of the knowledge base that is connected to the data source.")
		bedrockAgent_listKnowledgeBaseDocumentsCmd.Flags().String("max-results", "", "The maximum number of results to return in the response.")
		bedrockAgent_listKnowledgeBaseDocumentsCmd.Flags().String("next-token", "", "If the total number of results is greater than the `maxResults` value provided in the request, enter the token returned in the `nextToken` field in the response in this field to return the next batch of results.")
		bedrockAgent_listKnowledgeBaseDocumentsCmd.MarkFlagRequired("data-source-id")
		bedrockAgent_listKnowledgeBaseDocumentsCmd.MarkFlagRequired("knowledge-base-id")
	})
	bedrockAgentCmd.AddCommand(bedrockAgent_listKnowledgeBaseDocumentsCmd)
}
