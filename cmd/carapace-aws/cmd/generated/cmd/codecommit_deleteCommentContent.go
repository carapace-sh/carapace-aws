package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codecommit_deleteCommentContentCmd = &cobra.Command{
	Use:   "delete-comment-content",
	Short: "Deletes the content of a comment made on a change, file, or commit in a repository.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codecommit_deleteCommentContentCmd).Standalone()

	codecommit_deleteCommentContentCmd.Flags().String("comment-id", "", "The unique, system-generated ID of the comment.")
	codecommit_deleteCommentContentCmd.MarkFlagRequired("comment-id")
	codecommitCmd.AddCommand(codecommit_deleteCommentContentCmd)
}
