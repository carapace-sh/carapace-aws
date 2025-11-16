package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codecommit_mergeBranchesByThreeWayCmd = &cobra.Command{
	Use:   "merge-branches-by-three-way",
	Short: "Merges two specified branches using the three-way merge strategy.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codecommit_mergeBranchesByThreeWayCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codecommit_mergeBranchesByThreeWayCmd).Standalone()

		codecommit_mergeBranchesByThreeWayCmd.Flags().String("author-name", "", "The name of the author who created the commit.")
		codecommit_mergeBranchesByThreeWayCmd.Flags().String("commit-message", "", "The commit message to include in the commit information for the merge.")
		codecommit_mergeBranchesByThreeWayCmd.Flags().String("conflict-detail-level", "", "The level of conflict detail to use.")
		codecommit_mergeBranchesByThreeWayCmd.Flags().String("conflict-resolution", "", "If AUTOMERGE is the conflict resolution strategy, a list of inputs to use when resolving conflicts during a merge.")
		codecommit_mergeBranchesByThreeWayCmd.Flags().String("conflict-resolution-strategy", "", "Specifies which branch to use when resolving conflicts, or whether to attempt automatically merging two versions of a file.")
		codecommit_mergeBranchesByThreeWayCmd.Flags().String("destination-commit-specifier", "", "The branch, tag, HEAD, or other fully qualified reference used to identify a commit (for example, a branch name or a full commit ID).")
		codecommit_mergeBranchesByThreeWayCmd.Flags().String("email", "", "The email address of the person merging the branches.")
		codecommit_mergeBranchesByThreeWayCmd.Flags().String("keep-empty-folders", "", "If the commit contains deletions, whether to keep a folder or folder structure if the changes leave the folders empty.")
		codecommit_mergeBranchesByThreeWayCmd.Flags().String("repository-name", "", "The name of the repository where you want to merge two branches.")
		codecommit_mergeBranchesByThreeWayCmd.Flags().String("source-commit-specifier", "", "The branch, tag, HEAD, or other fully qualified reference used to identify a commit (for example, a branch name or a full commit ID).")
		codecommit_mergeBranchesByThreeWayCmd.Flags().String("target-branch", "", "The branch where the merge is applied.")
		codecommit_mergeBranchesByThreeWayCmd.MarkFlagRequired("destination-commit-specifier")
		codecommit_mergeBranchesByThreeWayCmd.MarkFlagRequired("repository-name")
		codecommit_mergeBranchesByThreeWayCmd.MarkFlagRequired("source-commit-specifier")
	})
	codecommitCmd.AddCommand(codecommit_mergeBranchesByThreeWayCmd)
}
