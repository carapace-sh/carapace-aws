package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workdocs_restoreDocumentVersionsCmd = &cobra.Command{
	Use:   "restore-document-versions",
	Short: "Recovers a deleted version of an Amazon WorkDocs document.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workdocs_restoreDocumentVersionsCmd).Standalone()

	workdocs_restoreDocumentVersionsCmd.Flags().String("authentication-token", "", "Amazon WorkDocs authentication token.")
	workdocs_restoreDocumentVersionsCmd.Flags().String("document-id", "", "The ID of the document.")
	workdocs_restoreDocumentVersionsCmd.MarkFlagRequired("document-id")
	workdocsCmd.AddCommand(workdocs_restoreDocumentVersionsCmd)
}
