package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workdocs_getDocumentVersionCmd = &cobra.Command{
	Use:   "get-document-version",
	Short: "Retrieves version metadata for the specified document.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workdocs_getDocumentVersionCmd).Standalone()

	workdocs_getDocumentVersionCmd.Flags().String("authentication-token", "", "Amazon WorkDocs authentication token.")
	workdocs_getDocumentVersionCmd.Flags().String("document-id", "", "The ID of the document.")
	workdocs_getDocumentVersionCmd.Flags().String("fields", "", "A comma-separated list of values.")
	workdocs_getDocumentVersionCmd.Flags().String("include-custom-metadata", "", "Set this to TRUE to include custom metadata in the response.")
	workdocs_getDocumentVersionCmd.Flags().String("version-id", "", "The version ID of the document.")
	workdocs_getDocumentVersionCmd.MarkFlagRequired("document-id")
	workdocs_getDocumentVersionCmd.MarkFlagRequired("version-id")
	workdocsCmd.AddCommand(workdocs_getDocumentVersionCmd)
}
