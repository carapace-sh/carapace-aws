package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workdocs_initiateDocumentVersionUploadCmd = &cobra.Command{
	Use:   "initiate-document-version-upload",
	Short: "Creates a new document object and version object.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workdocs_initiateDocumentVersionUploadCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(workdocs_initiateDocumentVersionUploadCmd).Standalone()

		workdocs_initiateDocumentVersionUploadCmd.Flags().String("authentication-token", "", "Amazon WorkDocs authentication token.")
		workdocs_initiateDocumentVersionUploadCmd.Flags().String("content-created-timestamp", "", "The timestamp when the content of the document was originally created.")
		workdocs_initiateDocumentVersionUploadCmd.Flags().String("content-modified-timestamp", "", "The timestamp when the content of the document was modified.")
		workdocs_initiateDocumentVersionUploadCmd.Flags().String("content-type", "", "The content type of the document.")
		workdocs_initiateDocumentVersionUploadCmd.Flags().String("document-size-in-bytes", "", "The size of the document, in bytes.")
		workdocs_initiateDocumentVersionUploadCmd.Flags().String("id", "", "The ID of the document.")
		workdocs_initiateDocumentVersionUploadCmd.Flags().String("name", "", "The name of the document.")
		workdocs_initiateDocumentVersionUploadCmd.Flags().String("parent-folder-id", "", "The ID of the parent folder.")
	})
	workdocsCmd.AddCommand(workdocs_initiateDocumentVersionUploadCmd)
}
