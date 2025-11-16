package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codecatalyst_listSourceRepositoryBranchesCmd = &cobra.Command{
	Use:   "list-source-repository-branches",
	Short: "Retrieves a list of branches in a specified source repository.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codecatalyst_listSourceRepositoryBranchesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codecatalyst_listSourceRepositoryBranchesCmd).Standalone()

		codecatalyst_listSourceRepositoryBranchesCmd.Flags().String("max-results", "", "The maximum number of results to show in a single call to this API.")
		codecatalyst_listSourceRepositoryBranchesCmd.Flags().String("next-token", "", "A token returned from a call to this API to indicate the next batch of results to return, if any.")
		codecatalyst_listSourceRepositoryBranchesCmd.Flags().String("project-name", "", "The name of the project in the space.")
		codecatalyst_listSourceRepositoryBranchesCmd.Flags().String("source-repository-name", "", "The name of the source repository.")
		codecatalyst_listSourceRepositoryBranchesCmd.Flags().String("space-name", "", "The name of the space.")
		codecatalyst_listSourceRepositoryBranchesCmd.MarkFlagRequired("project-name")
		codecatalyst_listSourceRepositoryBranchesCmd.MarkFlagRequired("source-repository-name")
		codecatalyst_listSourceRepositoryBranchesCmd.MarkFlagRequired("space-name")
	})
	codecatalystCmd.AddCommand(codecatalyst_listSourceRepositoryBranchesCmd)
}
