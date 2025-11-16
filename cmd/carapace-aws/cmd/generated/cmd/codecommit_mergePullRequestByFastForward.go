package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codecommit_mergePullRequestByFastForwardCmd = &cobra.Command{
	Use:   "merge-pull-request-by-fast-forward",
	Short: "Attempts to merge the source commit of a pull request into the specified destination branch for that pull request at the specified commit using the fast-forward merge strategy.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codecommit_mergePullRequestByFastForwardCmd).Standalone()

	codecommit_mergePullRequestByFastForwardCmd.Flags().String("pull-request-id", "", "The system-generated ID of the pull request.")
	codecommit_mergePullRequestByFastForwardCmd.Flags().String("repository-name", "", "The name of the repository where the pull request was created.")
	codecommit_mergePullRequestByFastForwardCmd.Flags().String("source-commit-id", "", "The full commit ID of the original or updated commit in the pull request source branch.")
	codecommit_mergePullRequestByFastForwardCmd.MarkFlagRequired("pull-request-id")
	codecommit_mergePullRequestByFastForwardCmd.MarkFlagRequired("repository-name")
	codecommitCmd.AddCommand(codecommit_mergePullRequestByFastForwardCmd)
}
