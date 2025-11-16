package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codecommit_getCommentsForPullRequestCmd = &cobra.Command{
	Use:   "get-comments-for-pull-request",
	Short: "Returns comments made on a pull request.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codecommit_getCommentsForPullRequestCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codecommit_getCommentsForPullRequestCmd).Standalone()

		codecommit_getCommentsForPullRequestCmd.Flags().String("after-commit-id", "", "The full commit ID of the commit in the source branch that was the tip of the branch at the time the comment was made.")
		codecommit_getCommentsForPullRequestCmd.Flags().String("before-commit-id", "", "The full commit ID of the commit in the destination branch that was the tip of the branch at the time the pull request was created.")
		codecommit_getCommentsForPullRequestCmd.Flags().String("max-results", "", "A non-zero, non-negative integer used to limit the number of returned results.")
		codecommit_getCommentsForPullRequestCmd.Flags().String("next-token", "", "An enumeration token that, when provided in a request, returns the next batch of the results.")
		codecommit_getCommentsForPullRequestCmd.Flags().String("pull-request-id", "", "The system-generated ID of the pull request.")
		codecommit_getCommentsForPullRequestCmd.Flags().String("repository-name", "", "The name of the repository that contains the pull request.")
		codecommit_getCommentsForPullRequestCmd.MarkFlagRequired("pull-request-id")
	})
	codecommitCmd.AddCommand(codecommit_getCommentsForPullRequestCmd)
}
