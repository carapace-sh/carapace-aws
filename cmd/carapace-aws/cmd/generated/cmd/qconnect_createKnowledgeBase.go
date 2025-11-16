package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qconnect_createKnowledgeBaseCmd = &cobra.Command{
	Use:   "create-knowledge-base",
	Short: "Creates a knowledge base.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qconnect_createKnowledgeBaseCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(qconnect_createKnowledgeBaseCmd).Standalone()

		qconnect_createKnowledgeBaseCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
		qconnect_createKnowledgeBaseCmd.Flags().String("description", "", "The description.")
		qconnect_createKnowledgeBaseCmd.Flags().String("knowledge-base-type", "", "The type of knowledge base.")
		qconnect_createKnowledgeBaseCmd.Flags().String("name", "", "The name of the knowledge base.")
		qconnect_createKnowledgeBaseCmd.Flags().String("rendering-configuration", "", "Information about how to render the content.")
		qconnect_createKnowledgeBaseCmd.Flags().String("server-side-encryption-configuration", "", "The configuration information for the customer managed key used for encryption.")
		qconnect_createKnowledgeBaseCmd.Flags().String("source-configuration", "", "The source of the knowledge base content.")
		qconnect_createKnowledgeBaseCmd.Flags().String("tags", "", "The tags used to organize, track, or control access for this resource.")
		qconnect_createKnowledgeBaseCmd.Flags().String("vector-ingestion-configuration", "", "Contains details about how to ingest the documents in a data source.")
		qconnect_createKnowledgeBaseCmd.MarkFlagRequired("knowledge-base-type")
		qconnect_createKnowledgeBaseCmd.MarkFlagRequired("name")
	})
	qconnectCmd.AddCommand(qconnect_createKnowledgeBaseCmd)
}
