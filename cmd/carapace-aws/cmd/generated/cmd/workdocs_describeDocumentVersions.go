package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workdocs_describeDocumentVersionsCmd = &cobra.Command{
	Use:   "describe-document-versions",
	Short: "Retrieves the document versions for the specified document.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workdocs_describeDocumentVersionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(workdocs_describeDocumentVersionsCmd).Standalone()

		workdocs_describeDocumentVersionsCmd.Flags().String("authentication-token", "", "Amazon WorkDocs authentication token.")
		workdocs_describeDocumentVersionsCmd.Flags().String("document-id", "", "The ID of the document.")
		workdocs_describeDocumentVersionsCmd.Flags().String("fields", "", "Specify \"SOURCE\" to include initialized versions and a URL for the source document.")
		workdocs_describeDocumentVersionsCmd.Flags().String("include", "", "A comma-separated list of values.")
		workdocs_describeDocumentVersionsCmd.Flags().String("limit", "", "The maximum number of versions to return with this call.")
		workdocs_describeDocumentVersionsCmd.Flags().String("marker", "", "The marker for the next set of results.")
		workdocs_describeDocumentVersionsCmd.MarkFlagRequired("document-id")
	})
	workdocsCmd.AddCommand(workdocs_describeDocumentVersionsCmd)
}
