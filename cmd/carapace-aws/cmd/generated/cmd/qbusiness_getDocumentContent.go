package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qbusiness_getDocumentContentCmd = &cobra.Command{
	Use:   "get-document-content",
	Short: "Retrieves the content of a document that was ingested into Amazon Q Business.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qbusiness_getDocumentContentCmd).Standalone()

	qbusiness_getDocumentContentCmd.Flags().String("application-id", "", "The unique identifier of the Amazon Q Business application containing the document.")
	qbusiness_getDocumentContentCmd.Flags().String("data-source-id", "", "The identifier of the data source from which the document was ingested.")
	qbusiness_getDocumentContentCmd.Flags().String("document-id", "", "The unique identifier of the document that is indexed via BatchPutDocument API or file-upload or connector sync.")
	qbusiness_getDocumentContentCmd.Flags().String("index-id", "", "The identifier of the index where documents are indexed.")
	qbusiness_getDocumentContentCmd.Flags().String("output-format", "", "Document outputFormat.")
	qbusiness_getDocumentContentCmd.MarkFlagRequired("application-id")
	qbusiness_getDocumentContentCmd.MarkFlagRequired("document-id")
	qbusiness_getDocumentContentCmd.MarkFlagRequired("index-id")
	qbusinessCmd.AddCommand(qbusiness_getDocumentContentCmd)
}
