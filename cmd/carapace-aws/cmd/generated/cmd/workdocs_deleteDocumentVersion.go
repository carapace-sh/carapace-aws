package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workdocs_deleteDocumentVersionCmd = &cobra.Command{
	Use:   "delete-document-version",
	Short: "Deletes a specific version of a document.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workdocs_deleteDocumentVersionCmd).Standalone()

	workdocs_deleteDocumentVersionCmd.Flags().String("authentication-token", "", "Amazon WorkDocs authentication token.")
	workdocs_deleteDocumentVersionCmd.Flags().String("delete-prior-versions", "", "Deletes all versions of a document prior to the current version.")
	workdocs_deleteDocumentVersionCmd.Flags().String("document-id", "", "The ID of the document associated with the version being deleted.")
	workdocs_deleteDocumentVersionCmd.Flags().String("version-id", "", "The ID of the version being deleted.")
	workdocs_deleteDocumentVersionCmd.MarkFlagRequired("delete-prior-versions")
	workdocs_deleteDocumentVersionCmd.MarkFlagRequired("document-id")
	workdocs_deleteDocumentVersionCmd.MarkFlagRequired("version-id")
	workdocsCmd.AddCommand(workdocs_deleteDocumentVersionCmd)
}
