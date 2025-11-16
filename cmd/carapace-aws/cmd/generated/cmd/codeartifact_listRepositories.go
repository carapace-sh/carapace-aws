package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codeartifact_listRepositoriesCmd = &cobra.Command{
	Use:   "list-repositories",
	Short: "Returns a list of [RepositorySummary](https://docs.aws.amazon.com/codeartifact/latest/APIReference/API_RepositorySummary.html) objects.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codeartifact_listRepositoriesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codeartifact_listRepositoriesCmd).Standalone()

		codeartifact_listRepositoriesCmd.Flags().String("max-results", "", "The maximum number of results to return per page.")
		codeartifact_listRepositoriesCmd.Flags().String("next-token", "", "The token for the next set of results.")
		codeartifact_listRepositoriesCmd.Flags().String("repository-prefix", "", "A prefix used to filter returned repositories.")
	})
	codeartifactCmd.AddCommand(codeartifact_listRepositoriesCmd)
}
