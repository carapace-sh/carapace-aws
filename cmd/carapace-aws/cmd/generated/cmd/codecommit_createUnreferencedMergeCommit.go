package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codecommit_createUnreferencedMergeCommitCmd = &cobra.Command{
	Use:   "create-unreferenced-merge-commit",
	Short: "Creates an unreferenced commit that represents the result of merging two branches using a specified merge strategy.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codecommit_createUnreferencedMergeCommitCmd).Standalone()

	codecommit_createUnreferencedMergeCommitCmd.Flags().String("author-name", "", "The name of the author who created the unreferenced commit.")
	codecommit_createUnreferencedMergeCommitCmd.Flags().String("commit-message", "", "The commit message for the unreferenced commit.")
	codecommit_createUnreferencedMergeCommitCmd.Flags().String("conflict-detail-level", "", "The level of conflict detail to use.")
	codecommit_createUnreferencedMergeCommitCmd.Flags().String("conflict-resolution", "", "If AUTOMERGE is the conflict resolution strategy, a list of inputs to use when resolving conflicts during a merge.")
	codecommit_createUnreferencedMergeCommitCmd.Flags().String("conflict-resolution-strategy", "", "Specifies which branch to use when resolving conflicts, or whether to attempt automatically merging two versions of a file.")
	codecommit_createUnreferencedMergeCommitCmd.Flags().String("destination-commit-specifier", "", "The branch, tag, HEAD, or other fully qualified reference used to identify a commit (for example, a branch name or a full commit ID).")
	codecommit_createUnreferencedMergeCommitCmd.Flags().String("email", "", "The email address for the person who created the unreferenced commit.")
	codecommit_createUnreferencedMergeCommitCmd.Flags().String("keep-empty-folders", "", "If the commit contains deletions, whether to keep a folder or folder structure if the changes leave the folders empty.")
	codecommit_createUnreferencedMergeCommitCmd.Flags().String("merge-option", "", "The merge option or strategy you want to use to merge the code.")
	codecommit_createUnreferencedMergeCommitCmd.Flags().String("repository-name", "", "The name of the repository where you want to create the unreferenced merge commit.")
	codecommit_createUnreferencedMergeCommitCmd.Flags().String("source-commit-specifier", "", "The branch, tag, HEAD, or other fully qualified reference used to identify a commit (for example, a branch name or a full commit ID).")
	codecommit_createUnreferencedMergeCommitCmd.MarkFlagRequired("destination-commit-specifier")
	codecommit_createUnreferencedMergeCommitCmd.MarkFlagRequired("merge-option")
	codecommit_createUnreferencedMergeCommitCmd.MarkFlagRequired("repository-name")
	codecommit_createUnreferencedMergeCommitCmd.MarkFlagRequired("source-commit-specifier")
	codecommitCmd.AddCommand(codecommit_createUnreferencedMergeCommitCmd)
}
