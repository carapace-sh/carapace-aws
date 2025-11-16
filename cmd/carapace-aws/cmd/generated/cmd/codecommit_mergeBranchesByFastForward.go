package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codecommit_mergeBranchesByFastForwardCmd = &cobra.Command{
	Use:   "merge-branches-by-fast-forward",
	Short: "Merges two branches using the fast-forward merge strategy.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codecommit_mergeBranchesByFastForwardCmd).Standalone()

	codecommit_mergeBranchesByFastForwardCmd.Flags().String("destination-commit-specifier", "", "The branch, tag, HEAD, or other fully qualified reference used to identify a commit (for example, a branch name or a full commit ID).")
	codecommit_mergeBranchesByFastForwardCmd.Flags().String("repository-name", "", "The name of the repository where you want to merge two branches.")
	codecommit_mergeBranchesByFastForwardCmd.Flags().String("source-commit-specifier", "", "The branch, tag, HEAD, or other fully qualified reference used to identify a commit (for example, a branch name or a full commit ID).")
	codecommit_mergeBranchesByFastForwardCmd.Flags().String("target-branch", "", "The branch where the merge is applied.")
	codecommit_mergeBranchesByFastForwardCmd.MarkFlagRequired("destination-commit-specifier")
	codecommit_mergeBranchesByFastForwardCmd.MarkFlagRequired("repository-name")
	codecommit_mergeBranchesByFastForwardCmd.MarkFlagRequired("source-commit-specifier")
	codecommitCmd.AddCommand(codecommit_mergeBranchesByFastForwardCmd)
}
