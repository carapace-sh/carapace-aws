package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codecommit_getMergeConflictsCmd = &cobra.Command{
	Use:   "get-merge-conflicts",
	Short: "Returns information about merge conflicts between the before and after commit IDs for a pull request in a repository.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codecommit_getMergeConflictsCmd).Standalone()

	codecommit_getMergeConflictsCmd.Flags().String("conflict-detail-level", "", "The level of conflict detail to use.")
	codecommit_getMergeConflictsCmd.Flags().String("conflict-resolution-strategy", "", "Specifies which branch to use when resolving conflicts, or whether to attempt automatically merging two versions of a file.")
	codecommit_getMergeConflictsCmd.Flags().String("destination-commit-specifier", "", "The branch, tag, HEAD, or other fully qualified reference used to identify a commit (for example, a branch name or a full commit ID).")
	codecommit_getMergeConflictsCmd.Flags().String("max-conflict-files", "", "The maximum number of files to include in the output.")
	codecommit_getMergeConflictsCmd.Flags().String("merge-option", "", "The merge option or strategy you want to use to merge the code.")
	codecommit_getMergeConflictsCmd.Flags().String("next-token", "", "An enumeration token that, when provided in a request, returns the next batch of the results.")
	codecommit_getMergeConflictsCmd.Flags().String("repository-name", "", "The name of the repository where the pull request was created.")
	codecommit_getMergeConflictsCmd.Flags().String("source-commit-specifier", "", "The branch, tag, HEAD, or other fully qualified reference used to identify a commit (for example, a branch name or a full commit ID).")
	codecommit_getMergeConflictsCmd.MarkFlagRequired("destination-commit-specifier")
	codecommit_getMergeConflictsCmd.MarkFlagRequired("merge-option")
	codecommit_getMergeConflictsCmd.MarkFlagRequired("repository-name")
	codecommit_getMergeConflictsCmd.MarkFlagRequired("source-commit-specifier")
	codecommitCmd.AddCommand(codecommit_getMergeConflictsCmd)
}
