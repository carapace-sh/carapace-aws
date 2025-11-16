package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codecommit_updateCommentCmd = &cobra.Command{
	Use:   "update-comment",
	Short: "Replaces the contents of a comment.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codecommit_updateCommentCmd).Standalone()

	codecommit_updateCommentCmd.Flags().String("comment-id", "", "The system-generated ID of the comment you want to update.")
	codecommit_updateCommentCmd.Flags().String("content", "", "The updated content to replace the existing content of the comment.")
	codecommit_updateCommentCmd.MarkFlagRequired("comment-id")
	codecommit_updateCommentCmd.MarkFlagRequired("content")
	codecommitCmd.AddCommand(codecommit_updateCommentCmd)
}
