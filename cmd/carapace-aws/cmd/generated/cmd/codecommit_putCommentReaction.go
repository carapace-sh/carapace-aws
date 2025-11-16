package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codecommit_putCommentReactionCmd = &cobra.Command{
	Use:   "put-comment-reaction",
	Short: "Adds or updates a reaction to a specified comment for the user whose identity is used to make the request.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codecommit_putCommentReactionCmd).Standalone()

	codecommit_putCommentReactionCmd.Flags().String("comment-id", "", "The ID of the comment to which you want to add or update a reaction.")
	codecommit_putCommentReactionCmd.Flags().String("reaction-value", "", "The emoji reaction you want to add or update.")
	codecommit_putCommentReactionCmd.MarkFlagRequired("comment-id")
	codecommit_putCommentReactionCmd.MarkFlagRequired("reaction-value")
	codecommitCmd.AddCommand(codecommit_putCommentReactionCmd)
}
