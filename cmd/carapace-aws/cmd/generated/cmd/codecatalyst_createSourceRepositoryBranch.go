package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codecatalyst_createSourceRepositoryBranchCmd = &cobra.Command{
	Use:   "create-source-repository-branch",
	Short: "Creates a branch in a specified source repository in Amazon CodeCatalyst.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codecatalyst_createSourceRepositoryBranchCmd).Standalone()

	codecatalyst_createSourceRepositoryBranchCmd.Flags().String("head-commit-id", "", "The commit ID in an existing branch from which you want to create the new branch.")
	codecatalyst_createSourceRepositoryBranchCmd.Flags().String("name", "", "The name for the branch you're creating.")
	codecatalyst_createSourceRepositoryBranchCmd.Flags().String("project-name", "", "The name of the project in the space.")
	codecatalyst_createSourceRepositoryBranchCmd.Flags().String("source-repository-name", "", "The name of the repository where you want to create a branch.")
	codecatalyst_createSourceRepositoryBranchCmd.Flags().String("space-name", "", "The name of the space.")
	codecatalyst_createSourceRepositoryBranchCmd.MarkFlagRequired("name")
	codecatalyst_createSourceRepositoryBranchCmd.MarkFlagRequired("project-name")
	codecatalyst_createSourceRepositoryBranchCmd.MarkFlagRequired("source-repository-name")
	codecatalyst_createSourceRepositoryBranchCmd.MarkFlagRequired("space-name")
	codecatalystCmd.AddCommand(codecatalyst_createSourceRepositoryBranchCmd)
}
