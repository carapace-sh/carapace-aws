package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codecommit_updateDefaultBranchCmd = &cobra.Command{
	Use:   "update-default-branch",
	Short: "Sets or changes the default branch name for the specified repository.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codecommit_updateDefaultBranchCmd).Standalone()

	codecommit_updateDefaultBranchCmd.Flags().String("default-branch-name", "", "The name of the branch to set as the default branch.")
	codecommit_updateDefaultBranchCmd.Flags().String("repository-name", "", "The name of the repository for which you want to set or change the default branch.")
	codecommit_updateDefaultBranchCmd.MarkFlagRequired("default-branch-name")
	codecommit_updateDefaultBranchCmd.MarkFlagRequired("repository-name")
	codecommitCmd.AddCommand(codecommit_updateDefaultBranchCmd)
}
