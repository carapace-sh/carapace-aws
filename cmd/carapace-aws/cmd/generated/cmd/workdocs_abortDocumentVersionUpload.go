package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workdocs_abortDocumentVersionUploadCmd = &cobra.Command{
	Use:   "abort-document-version-upload",
	Short: "Aborts the upload of the specified document version that was previously initiated by [InitiateDocumentVersionUpload]().",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workdocs_abortDocumentVersionUploadCmd).Standalone()

	workdocs_abortDocumentVersionUploadCmd.Flags().String("authentication-token", "", "Amazon WorkDocs authentication token.")
	workdocs_abortDocumentVersionUploadCmd.Flags().String("document-id", "", "The ID of the document.")
	workdocs_abortDocumentVersionUploadCmd.Flags().String("version-id", "", "The ID of the version.")
	workdocs_abortDocumentVersionUploadCmd.MarkFlagRequired("document-id")
	workdocs_abortDocumentVersionUploadCmd.MarkFlagRequired("version-id")
	workdocsCmd.AddCommand(workdocs_abortDocumentVersionUploadCmd)
}
