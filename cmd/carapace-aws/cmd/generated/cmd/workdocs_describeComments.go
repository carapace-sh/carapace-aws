package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workdocs_describeCommentsCmd = &cobra.Command{
	Use:   "describe-comments",
	Short: "List all the comments for the specified document version.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workdocs_describeCommentsCmd).Standalone()

	workdocs_describeCommentsCmd.Flags().String("authentication-token", "", "Amazon WorkDocs authentication token.")
	workdocs_describeCommentsCmd.Flags().String("document-id", "", "The ID of the document.")
	workdocs_describeCommentsCmd.Flags().String("limit", "", "The maximum number of items to return.")
	workdocs_describeCommentsCmd.Flags().String("marker", "", "The marker for the next set of results.")
	workdocs_describeCommentsCmd.Flags().String("version-id", "", "The ID of the document version.")
	workdocs_describeCommentsCmd.MarkFlagRequired("document-id")
	workdocs_describeCommentsCmd.MarkFlagRequired("version-id")
	workdocsCmd.AddCommand(workdocs_describeCommentsCmd)
}
