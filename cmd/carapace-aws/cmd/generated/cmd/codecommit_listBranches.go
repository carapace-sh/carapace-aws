package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codecommit_listBranchesCmd = &cobra.Command{
	Use:   "list-branches",
	Short: "Gets information about one or more branches in a repository.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codecommit_listBranchesCmd).Standalone()

	codecommit_listBranchesCmd.Flags().String("next-token", "", "An enumeration token that allows the operation to batch the results.")
	codecommit_listBranchesCmd.Flags().String("repository-name", "", "The name of the repository that contains the branches.")
	codecommit_listBranchesCmd.MarkFlagRequired("repository-name")
	codecommitCmd.AddCommand(codecommit_listBranchesCmd)
}
