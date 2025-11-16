package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codecommit_batchDescribeMergeConflictsCmd = &cobra.Command{
	Use:   "batch-describe-merge-conflicts",
	Short: "Returns information about one or more merge conflicts in the attempted merge of two commit specifiers using the squash or three-way merge strategy.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codecommit_batchDescribeMergeConflictsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codecommit_batchDescribeMergeConflictsCmd).Standalone()

		codecommit_batchDescribeMergeConflictsCmd.Flags().String("conflict-detail-level", "", "The level of conflict detail to use.")
		codecommit_batchDescribeMergeConflictsCmd.Flags().String("conflict-resolution-strategy", "", "Specifies which branch to use when resolving conflicts, or whether to attempt automatically merging two versions of a file.")
		codecommit_batchDescribeMergeConflictsCmd.Flags().String("destination-commit-specifier", "", "The branch, tag, HEAD, or other fully qualified reference used to identify a commit (for example, a branch name or a full commit ID).")
		codecommit_batchDescribeMergeConflictsCmd.Flags().String("file-paths", "", "The path of the target files used to describe the conflicts.")
		codecommit_batchDescribeMergeConflictsCmd.Flags().String("max-conflict-files", "", "The maximum number of files to include in the output.")
		codecommit_batchDescribeMergeConflictsCmd.Flags().String("max-merge-hunks", "", "The maximum number of merge hunks to include in the output.")
		codecommit_batchDescribeMergeConflictsCmd.Flags().String("merge-option", "", "The merge option or strategy you want to use to merge the code.")
		codecommit_batchDescribeMergeConflictsCmd.Flags().String("next-token", "", "An enumeration token that, when provided in a request, returns the next batch of the results.")
		codecommit_batchDescribeMergeConflictsCmd.Flags().String("repository-name", "", "The name of the repository that contains the merge conflicts you want to review.")
		codecommit_batchDescribeMergeConflictsCmd.Flags().String("source-commit-specifier", "", "The branch, tag, HEAD, or other fully qualified reference used to identify a commit (for example, a branch name or a full commit ID).")
		codecommit_batchDescribeMergeConflictsCmd.MarkFlagRequired("destination-commit-specifier")
		codecommit_batchDescribeMergeConflictsCmd.MarkFlagRequired("merge-option")
		codecommit_batchDescribeMergeConflictsCmd.MarkFlagRequired("repository-name")
		codecommit_batchDescribeMergeConflictsCmd.MarkFlagRequired("source-commit-specifier")
	})
	codecommitCmd.AddCommand(codecommit_batchDescribeMergeConflictsCmd)
}
