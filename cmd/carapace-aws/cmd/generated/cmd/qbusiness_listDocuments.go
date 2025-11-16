package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qbusiness_listDocumentsCmd = &cobra.Command{
	Use:   "list-documents",
	Short: "A list of documents attached to an index.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qbusiness_listDocumentsCmd).Standalone()

	qbusiness_listDocumentsCmd.Flags().String("application-id", "", "The identifier of the application id the documents are attached to.")
	qbusiness_listDocumentsCmd.Flags().String("data-source-ids", "", "The identifier of the data sources the documents are attached to.")
	qbusiness_listDocumentsCmd.Flags().String("index-id", "", "The identifier of the index the documents are attached to.")
	qbusiness_listDocumentsCmd.Flags().String("max-results", "", "The maximum number of documents to return.")
	qbusiness_listDocumentsCmd.Flags().String("next-token", "", "If the `maxResults` response was incomplete because there is more data to retrieve, Amazon Q Business returns a pagination token in the response.")
	qbusiness_listDocumentsCmd.MarkFlagRequired("application-id")
	qbusiness_listDocumentsCmd.MarkFlagRequired("index-id")
	qbusinessCmd.AddCommand(qbusiness_listDocumentsCmd)
}
