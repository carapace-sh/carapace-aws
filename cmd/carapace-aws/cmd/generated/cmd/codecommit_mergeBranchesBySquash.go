package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codecommit_mergeBranchesBySquashCmd = &cobra.Command{
	Use:   "merge-branches-by-squash",
	Short: "Merges two branches using the squash merge strategy.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codecommit_mergeBranchesBySquashCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codecommit_mergeBranchesBySquashCmd).Standalone()

		codecommit_mergeBranchesBySquashCmd.Flags().String("author-name", "", "The name of the author who created the commit.")
		codecommit_mergeBranchesBySquashCmd.Flags().String("commit-message", "", "The commit message for the merge.")
		codecommit_mergeBranchesBySquashCmd.Flags().String("conflict-detail-level", "", "The level of conflict detail to use.")
		codecommit_mergeBranchesBySquashCmd.Flags().String("conflict-resolution", "", "If AUTOMERGE is the conflict resolution strategy, a list of inputs to use when resolving conflicts during a merge.")
		codecommit_mergeBranchesBySquashCmd.Flags().String("conflict-resolution-strategy", "", "Specifies which branch to use when resolving conflicts, or whether to attempt automatically merging two versions of a file.")
		codecommit_mergeBranchesBySquashCmd.Flags().String("destination-commit-specifier", "", "The branch, tag, HEAD, or other fully qualified reference used to identify a commit (for example, a branch name or a full commit ID).")
		codecommit_mergeBranchesBySquashCmd.Flags().String("email", "", "The email address of the person merging the branches.")
		codecommit_mergeBranchesBySquashCmd.Flags().String("keep-empty-folders", "", "If the commit contains deletions, whether to keep a folder or folder structure if the changes leave the folders empty.")
		codecommit_mergeBranchesBySquashCmd.Flags().String("repository-name", "", "The name of the repository where you want to merge two branches.")
		codecommit_mergeBranchesBySquashCmd.Flags().String("source-commit-specifier", "", "The branch, tag, HEAD, or other fully qualified reference used to identify a commit (for example, a branch name or a full commit ID).")
		codecommit_mergeBranchesBySquashCmd.Flags().String("target-branch", "", "The branch where the merge is applied.")
		codecommit_mergeBranchesBySquashCmd.MarkFlagRequired("destination-commit-specifier")
		codecommit_mergeBranchesBySquashCmd.MarkFlagRequired("repository-name")
		codecommit_mergeBranchesBySquashCmd.MarkFlagRequired("source-commit-specifier")
	})
	codecommitCmd.AddCommand(codecommit_mergeBranchesBySquashCmd)
}
