package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codecommit_createBranchCmd = &cobra.Command{
	Use:   "create-branch",
	Short: "Creates a branch in a repository and points the branch to a commit.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codecommit_createBranchCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codecommit_createBranchCmd).Standalone()

		codecommit_createBranchCmd.Flags().String("branch-name", "", "The name of the new branch to create.")
		codecommit_createBranchCmd.Flags().String("commit-id", "", "The ID of the commit to point the new branch to.")
		codecommit_createBranchCmd.Flags().String("repository-name", "", "The name of the repository in which you want to create the new branch.")
		codecommit_createBranchCmd.MarkFlagRequired("branch-name")
		codecommit_createBranchCmd.MarkFlagRequired("commit-id")
		codecommit_createBranchCmd.MarkFlagRequired("repository-name")
	})
	codecommitCmd.AddCommand(codecommit_createBranchCmd)
}
