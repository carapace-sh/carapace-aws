package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workdocs_createCommentCmd = &cobra.Command{
	Use:   "create-comment",
	Short: "Adds a new comment to the specified document version.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workdocs_createCommentCmd).Standalone()

	workdocs_createCommentCmd.Flags().String("authentication-token", "", "Amazon WorkDocs authentication token.")
	workdocs_createCommentCmd.Flags().String("document-id", "", "The ID of the document.")
	workdocs_createCommentCmd.Flags().String("notify-collaborators", "", "Set this parameter to TRUE to send an email out to the document collaborators after the comment is created.")
	workdocs_createCommentCmd.Flags().String("parent-id", "", "The ID of the parent comment.")
	workdocs_createCommentCmd.Flags().String("text", "", "The text of the comment.")
	workdocs_createCommentCmd.Flags().String("thread-id", "", "The ID of the root comment in the thread.")
	workdocs_createCommentCmd.Flags().String("version-id", "", "The ID of the document version.")
	workdocs_createCommentCmd.Flags().String("visibility", "", "The visibility of the comment.")
	workdocs_createCommentCmd.MarkFlagRequired("document-id")
	workdocs_createCommentCmd.MarkFlagRequired("text")
	workdocs_createCommentCmd.MarkFlagRequired("version-id")
	workdocsCmd.AddCommand(workdocs_createCommentCmd)
}
