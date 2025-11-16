package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codecommit_postCommentForPullRequestCmd = &cobra.Command{
	Use:   "post-comment-for-pull-request",
	Short: "Posts a comment on a pull request.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codecommit_postCommentForPullRequestCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codecommit_postCommentForPullRequestCmd).Standalone()

		codecommit_postCommentForPullRequestCmd.Flags().String("after-commit-id", "", "The full commit ID of the commit in the source branch that is the current tip of the branch for the pull request when you post the comment.")
		codecommit_postCommentForPullRequestCmd.Flags().String("before-commit-id", "", "The full commit ID of the commit in the destination branch that was the tip of the branch at the time the pull request was created.")
		codecommit_postCommentForPullRequestCmd.Flags().String("client-request-token", "", "A unique, client-generated idempotency token that, when provided in a request, ensures the request cannot be repeated with a changed parameter.")
		codecommit_postCommentForPullRequestCmd.Flags().String("content", "", "The content of your comment on the change.")
		codecommit_postCommentForPullRequestCmd.Flags().String("location", "", "The location of the change where you want to post your comment.")
		codecommit_postCommentForPullRequestCmd.Flags().String("pull-request-id", "", "The system-generated ID of the pull request.")
		codecommit_postCommentForPullRequestCmd.Flags().String("repository-name", "", "The name of the repository where you want to post a comment on a pull request.")
		codecommit_postCommentForPullRequestCmd.MarkFlagRequired("after-commit-id")
		codecommit_postCommentForPullRequestCmd.MarkFlagRequired("before-commit-id")
		codecommit_postCommentForPullRequestCmd.MarkFlagRequired("content")
		codecommit_postCommentForPullRequestCmd.MarkFlagRequired("pull-request-id")
		codecommit_postCommentForPullRequestCmd.MarkFlagRequired("repository-name")
	})
	codecommitCmd.AddCommand(codecommit_postCommentForPullRequestCmd)
}
