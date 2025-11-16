package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codecommit_getMergeOptionsCmd = &cobra.Command{
	Use:   "get-merge-options",
	Short: "Returns information about the merge options available for merging two specified branches.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codecommit_getMergeOptionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codecommit_getMergeOptionsCmd).Standalone()

		codecommit_getMergeOptionsCmd.Flags().String("conflict-detail-level", "", "The level of conflict detail to use.")
		codecommit_getMergeOptionsCmd.Flags().String("conflict-resolution-strategy", "", "Specifies which branch to use when resolving conflicts, or whether to attempt automatically merging two versions of a file.")
		codecommit_getMergeOptionsCmd.Flags().String("destination-commit-specifier", "", "The branch, tag, HEAD, or other fully qualified reference used to identify a commit (for example, a branch name or a full commit ID).")
		codecommit_getMergeOptionsCmd.Flags().String("repository-name", "", "The name of the repository that contains the commits about which you want to get merge options.")
		codecommit_getMergeOptionsCmd.Flags().String("source-commit-specifier", "", "The branch, tag, HEAD, or other fully qualified reference used to identify a commit (for example, a branch name or a full commit ID).")
		codecommit_getMergeOptionsCmd.MarkFlagRequired("destination-commit-specifier")
		codecommit_getMergeOptionsCmd.MarkFlagRequired("repository-name")
		codecommit_getMergeOptionsCmd.MarkFlagRequired("source-commit-specifier")
	})
	codecommitCmd.AddCommand(codecommit_getMergeOptionsCmd)
}
