package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codecommit_describeMergeConflictsCmd = &cobra.Command{
	Use:   "describe-merge-conflicts",
	Short: "Returns information about one or more merge conflicts in the attempted merge of two commit specifiers using the squash or three-way merge strategy.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codecommit_describeMergeConflictsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codecommit_describeMergeConflictsCmd).Standalone()

		codecommit_describeMergeConflictsCmd.Flags().String("conflict-detail-level", "", "The level of conflict detail to use.")
		codecommit_describeMergeConflictsCmd.Flags().String("conflict-resolution-strategy", "", "Specifies which branch to use when resolving conflicts, or whether to attempt automatically merging two versions of a file.")
		codecommit_describeMergeConflictsCmd.Flags().String("destination-commit-specifier", "", "The branch, tag, HEAD, or other fully qualified reference used to identify a commit (for example, a branch name or a full commit ID).")
		codecommit_describeMergeConflictsCmd.Flags().String("file-path", "", "The path of the target files used to describe the conflicts.")
		codecommit_describeMergeConflictsCmd.Flags().String("max-merge-hunks", "", "The maximum number of merge hunks to include in the output.")
		codecommit_describeMergeConflictsCmd.Flags().String("merge-option", "", "The merge option or strategy you want to use to merge the code.")
		codecommit_describeMergeConflictsCmd.Flags().String("next-token", "", "An enumeration token that, when provided in a request, returns the next batch of the results.")
		codecommit_describeMergeConflictsCmd.Flags().String("repository-name", "", "The name of the repository where you want to get information about a merge conflict.")
		codecommit_describeMergeConflictsCmd.Flags().String("source-commit-specifier", "", "The branch, tag, HEAD, or other fully qualified reference used to identify a commit (for example, a branch name or a full commit ID).")
		codecommit_describeMergeConflictsCmd.MarkFlagRequired("destination-commit-specifier")
		codecommit_describeMergeConflictsCmd.MarkFlagRequired("file-path")
		codecommit_describeMergeConflictsCmd.MarkFlagRequired("merge-option")
		codecommit_describeMergeConflictsCmd.MarkFlagRequired("repository-name")
		codecommit_describeMergeConflictsCmd.MarkFlagRequired("source-commit-specifier")
	})
	codecommitCmd.AddCommand(codecommit_describeMergeConflictsCmd)
}
