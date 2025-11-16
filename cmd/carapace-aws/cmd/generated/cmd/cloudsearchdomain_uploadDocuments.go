package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudsearchdomain_uploadDocumentsCmd = &cobra.Command{
	Use:   "upload-documents",
	Short: "Posts a batch of documents to a search domain for indexing.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudsearchdomain_uploadDocumentsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloudsearchdomain_uploadDocumentsCmd).Standalone()

		cloudsearchdomain_uploadDocumentsCmd.Flags().String("content-type", "", "The format of the batch you are uploading.")
		cloudsearchdomain_uploadDocumentsCmd.Flags().String("documents", "", "A batch of documents formatted in JSON or HTML.")
		cloudsearchdomain_uploadDocumentsCmd.MarkFlagRequired("content-type")
		cloudsearchdomain_uploadDocumentsCmd.MarkFlagRequired("documents")
	})
	cloudsearchdomainCmd.AddCommand(cloudsearchdomain_uploadDocumentsCmd)
}
