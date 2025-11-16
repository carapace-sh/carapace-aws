package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workdocs_updateDocumentVersionCmd = &cobra.Command{
	Use:   "update-document-version",
	Short: "Changes the status of the document version to ACTIVE.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workdocs_updateDocumentVersionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(workdocs_updateDocumentVersionCmd).Standalone()

		workdocs_updateDocumentVersionCmd.Flags().String("authentication-token", "", "Amazon WorkDocs authentication token.")
		workdocs_updateDocumentVersionCmd.Flags().String("document-id", "", "The ID of the document.")
		workdocs_updateDocumentVersionCmd.Flags().String("version-id", "", "The version ID of the document.")
		workdocs_updateDocumentVersionCmd.Flags().String("version-status", "", "The status of the version.")
		workdocs_updateDocumentVersionCmd.MarkFlagRequired("document-id")
		workdocs_updateDocumentVersionCmd.MarkFlagRequired("version-id")
	})
	workdocsCmd.AddCommand(workdocs_updateDocumentVersionCmd)
}
