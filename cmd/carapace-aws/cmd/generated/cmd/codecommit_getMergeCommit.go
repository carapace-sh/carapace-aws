package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codecommit_getMergeCommitCmd = &cobra.Command{
	Use:   "get-merge-commit",
	Short: "Returns information about a specified merge commit.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codecommit_getMergeCommitCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codecommit_getMergeCommitCmd).Standalone()

		codecommit_getMergeCommitCmd.Flags().String("conflict-detail-level", "", "The level of conflict detail to use.")
		codecommit_getMergeCommitCmd.Flags().String("conflict-resolution-strategy", "", "Specifies which branch to use when resolving conflicts, or whether to attempt automatically merging two versions of a file.")
		codecommit_getMergeCommitCmd.Flags().String("destination-commit-specifier", "", "The branch, tag, HEAD, or other fully qualified reference used to identify a commit (for example, a branch name or a full commit ID).")
		codecommit_getMergeCommitCmd.Flags().String("repository-name", "", "The name of the repository that contains the merge commit about which you want to get information.")
		codecommit_getMergeCommitCmd.Flags().String("source-commit-specifier", "", "The branch, tag, HEAD, or other fully qualified reference used to identify a commit (for example, a branch name or a full commit ID).")
		codecommit_getMergeCommitCmd.MarkFlagRequired("destination-commit-specifier")
		codecommit_getMergeCommitCmd.MarkFlagRequired("repository-name")
		codecommit_getMergeCommitCmd.MarkFlagRequired("source-commit-specifier")
	})
	codecommitCmd.AddCommand(codecommit_getMergeCommitCmd)
}
