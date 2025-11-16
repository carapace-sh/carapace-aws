package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codecommit_deleteBranchCmd = &cobra.Command{
	Use:   "delete-branch",
	Short: "Deletes a branch from a repository, unless that branch is the default branch for the repository.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codecommit_deleteBranchCmd).Standalone()

	codecommit_deleteBranchCmd.Flags().String("branch-name", "", "The name of the branch to delete.")
	codecommit_deleteBranchCmd.Flags().String("repository-name", "", "The name of the repository that contains the branch to be deleted.")
	codecommit_deleteBranchCmd.MarkFlagRequired("branch-name")
	codecommit_deleteBranchCmd.MarkFlagRequired("repository-name")
	codecommitCmd.AddCommand(codecommit_deleteBranchCmd)
}
