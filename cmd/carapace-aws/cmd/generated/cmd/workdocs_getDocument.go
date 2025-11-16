package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workdocs_getDocumentCmd = &cobra.Command{
	Use:   "get-document",
	Short: "Retrieves details of a document.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workdocs_getDocumentCmd).Standalone()

	workdocs_getDocumentCmd.Flags().String("authentication-token", "", "Amazon WorkDocs authentication token.")
	workdocs_getDocumentCmd.Flags().String("document-id", "", "The ID of the document.")
	workdocs_getDocumentCmd.Flags().String("include-custom-metadata", "", "Set this to `TRUE` to include custom metadata in the response.")
	workdocs_getDocumentCmd.MarkFlagRequired("document-id")
	workdocsCmd.AddCommand(workdocs_getDocumentCmd)
}
