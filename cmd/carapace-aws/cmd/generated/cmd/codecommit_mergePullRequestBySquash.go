package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codecommit_mergePullRequestBySquashCmd = &cobra.Command{
	Use:   "merge-pull-request-by-squash",
	Short: "Attempts to merge the source commit of a pull request into the specified destination branch for that pull request at the specified commit using the squash merge strategy.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codecommit_mergePullRequestBySquashCmd).Standalone()

	codecommit_mergePullRequestBySquashCmd.Flags().String("author-name", "", "The name of the author who created the commit.")
	codecommit_mergePullRequestBySquashCmd.Flags().String("commit-message", "", "The commit message to include in the commit information for the merge.")
	codecommit_mergePullRequestBySquashCmd.Flags().String("conflict-detail-level", "", "The level of conflict detail to use.")
	codecommit_mergePullRequestBySquashCmd.Flags().String("conflict-resolution", "", "If AUTOMERGE is the conflict resolution strategy, a list of inputs to use when resolving conflicts during a merge.")
	codecommit_mergePullRequestBySquashCmd.Flags().String("conflict-resolution-strategy", "", "Specifies which branch to use when resolving conflicts, or whether to attempt automatically merging two versions of a file.")
	codecommit_mergePullRequestBySquashCmd.Flags().String("email", "", "The email address of the person merging the branches.")
	codecommit_mergePullRequestBySquashCmd.Flags().String("keep-empty-folders", "", "If the commit contains deletions, whether to keep a folder or folder structure if the changes leave the folders empty.")
	codecommit_mergePullRequestBySquashCmd.Flags().String("pull-request-id", "", "The system-generated ID of the pull request.")
	codecommit_mergePullRequestBySquashCmd.Flags().String("repository-name", "", "The name of the repository where the pull request was created.")
	codecommit_mergePullRequestBySquashCmd.Flags().String("source-commit-id", "", "The full commit ID of the original or updated commit in the pull request source branch.")
	codecommit_mergePullRequestBySquashCmd.MarkFlagRequired("pull-request-id")
	codecommit_mergePullRequestBySquashCmd.MarkFlagRequired("repository-name")
	codecommitCmd.AddCommand(codecommit_mergePullRequestBySquashCmd)
}
