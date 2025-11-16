package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workdocs_deleteCommentCmd = &cobra.Command{
	Use:   "delete-comment",
	Short: "Deletes the specified comment from the document version.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workdocs_deleteCommentCmd).Standalone()

	workdocs_deleteCommentCmd.Flags().String("authentication-token", "", "Amazon WorkDocs authentication token.")
	workdocs_deleteCommentCmd.Flags().String("comment-id", "", "The ID of the comment.")
	workdocs_deleteCommentCmd.Flags().String("document-id", "", "The ID of the document.")
	workdocs_deleteCommentCmd.Flags().String("version-id", "", "The ID of the document version.")
	workdocs_deleteCommentCmd.MarkFlagRequired("comment-id")
	workdocs_deleteCommentCmd.MarkFlagRequired("document-id")
	workdocs_deleteCommentCmd.MarkFlagRequired("version-id")
	workdocsCmd.AddCommand(workdocs_deleteCommentCmd)
}
