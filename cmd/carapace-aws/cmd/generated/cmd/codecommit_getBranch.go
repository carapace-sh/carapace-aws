package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codecommit_getBranchCmd = &cobra.Command{
	Use:   "get-branch",
	Short: "Returns information about a repository branch, including its name and the last commit ID.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codecommit_getBranchCmd).Standalone()

	codecommit_getBranchCmd.Flags().String("branch-name", "", "The name of the branch for which you want to retrieve information.")
	codecommit_getBranchCmd.Flags().String("repository-name", "", "The name of the repository that contains the branch for which you want to retrieve information.")
	codecommitCmd.AddCommand(codecommit_getBranchCmd)
}
