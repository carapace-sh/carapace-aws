package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workdocs_deleteDocumentCmd = &cobra.Command{
	Use:   "delete-document",
	Short: "Permanently deletes the specified document and its associated metadata.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workdocs_deleteDocumentCmd).Standalone()

	workdocs_deleteDocumentCmd.Flags().String("authentication-token", "", "Amazon WorkDocs authentication token.")
	workdocs_deleteDocumentCmd.Flags().String("document-id", "", "The ID of the document.")
	workdocs_deleteDocumentCmd.MarkFlagRequired("document-id")
	workdocsCmd.AddCommand(workdocs_deleteDocumentCmd)
}
