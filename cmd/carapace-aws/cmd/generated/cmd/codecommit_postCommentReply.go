package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codecommit_postCommentReplyCmd = &cobra.Command{
	Use:   "post-comment-reply",
	Short: "Posts a comment in reply to an existing comment on a comparison between commits or a pull request.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codecommit_postCommentReplyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codecommit_postCommentReplyCmd).Standalone()

		codecommit_postCommentReplyCmd.Flags().String("client-request-token", "", "A unique, client-generated idempotency token that, when provided in a request, ensures the request cannot be repeated with a changed parameter.")
		codecommit_postCommentReplyCmd.Flags().String("content", "", "The contents of your reply to a comment.")
		codecommit_postCommentReplyCmd.Flags().String("in-reply-to", "", "The system-generated ID of the comment to which you want to reply.")
		codecommit_postCommentReplyCmd.MarkFlagRequired("content")
		codecommit_postCommentReplyCmd.MarkFlagRequired("in-reply-to")
	})
	codecommitCmd.AddCommand(codecommit_postCommentReplyCmd)
}
