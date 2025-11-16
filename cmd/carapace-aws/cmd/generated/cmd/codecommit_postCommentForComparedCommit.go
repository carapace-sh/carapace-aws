package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codecommit_postCommentForComparedCommitCmd = &cobra.Command{
	Use:   "post-comment-for-compared-commit",
	Short: "Posts a comment on the comparison between two commits.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codecommit_postCommentForComparedCommitCmd).Standalone()

	codecommit_postCommentForComparedCommitCmd.Flags().String("after-commit-id", "", "To establish the directionality of the comparison, the full commit ID of the after commit.")
	codecommit_postCommentForComparedCommitCmd.Flags().String("before-commit-id", "", "To establish the directionality of the comparison, the full commit ID of the before commit.")
	codecommit_postCommentForComparedCommitCmd.Flags().String("client-request-token", "", "A unique, client-generated idempotency token that, when provided in a request, ensures the request cannot be repeated with a changed parameter.")
	codecommit_postCommentForComparedCommitCmd.Flags().String("content", "", "The content of the comment you want to make.")
	codecommit_postCommentForComparedCommitCmd.Flags().String("location", "", "The location of the comparison where you want to comment.")
	codecommit_postCommentForComparedCommitCmd.Flags().String("repository-name", "", "The name of the repository where you want to post a comment on the comparison between commits.")
	codecommit_postCommentForComparedCommitCmd.MarkFlagRequired("after-commit-id")
	codecommit_postCommentForComparedCommitCmd.MarkFlagRequired("content")
	codecommit_postCommentForComparedCommitCmd.MarkFlagRequired("repository-name")
	codecommitCmd.AddCommand(codecommit_postCommentForComparedCommitCmd)
}
