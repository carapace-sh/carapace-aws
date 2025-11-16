package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workdocs_updateDocumentCmd = &cobra.Command{
	Use:   "update-document",
	Short: "Updates the specified attributes of a document.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workdocs_updateDocumentCmd).Standalone()

	workdocs_updateDocumentCmd.Flags().String("authentication-token", "", "Amazon WorkDocs authentication token.")
	workdocs_updateDocumentCmd.Flags().String("document-id", "", "The ID of the document.")
	workdocs_updateDocumentCmd.Flags().String("name", "", "The name of the document.")
	workdocs_updateDocumentCmd.Flags().String("parent-folder-id", "", "The ID of the parent folder.")
	workdocs_updateDocumentCmd.Flags().String("resource-state", "", "The resource state of the document.")
	workdocs_updateDocumentCmd.MarkFlagRequired("document-id")
	workdocsCmd.AddCommand(workdocs_updateDocumentCmd)
}
