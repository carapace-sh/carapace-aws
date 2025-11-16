package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workdocs_getDocumentPathCmd = &cobra.Command{
	Use:   "get-document-path",
	Short: "Retrieves the path information (the hierarchy from the root folder) for the requested document.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workdocs_getDocumentPathCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(workdocs_getDocumentPathCmd).Standalone()

		workdocs_getDocumentPathCmd.Flags().String("authentication-token", "", "Amazon WorkDocs authentication token.")
		workdocs_getDocumentPathCmd.Flags().String("document-id", "", "The ID of the document.")
		workdocs_getDocumentPathCmd.Flags().String("fields", "", "A comma-separated list of values.")
		workdocs_getDocumentPathCmd.Flags().String("limit", "", "The maximum number of levels in the hierarchy to return.")
		workdocs_getDocumentPathCmd.Flags().String("marker", "", "This value is not supported.")
		workdocs_getDocumentPathCmd.MarkFlagRequired("document-id")
	})
	workdocsCmd.AddCommand(workdocs_getDocumentPathCmd)
}
