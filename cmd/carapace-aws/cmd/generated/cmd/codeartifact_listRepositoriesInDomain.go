package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codeartifact_listRepositoriesInDomainCmd = &cobra.Command{
	Use:   "list-repositories-in-domain",
	Short: "Returns a list of [RepositorySummary](https://docs.aws.amazon.com/codeartifact/latest/APIReference/API_RepositorySummary.html) objects.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codeartifact_listRepositoriesInDomainCmd).Standalone()

	codeartifact_listRepositoriesInDomainCmd.Flags().String("administrator-account", "", "Filter the list of repositories to only include those that are managed by the Amazon Web Services account ID.")
	codeartifact_listRepositoriesInDomainCmd.Flags().String("domain", "", "The name of the domain that contains the returned list of repositories.")
	codeartifact_listRepositoriesInDomainCmd.Flags().String("domain-owner", "", "The 12-digit account number of the Amazon Web Services account that owns the domain.")
	codeartifact_listRepositoriesInDomainCmd.Flags().String("max-results", "", "The maximum number of results to return per page.")
	codeartifact_listRepositoriesInDomainCmd.Flags().String("next-token", "", "The token for the next set of results.")
	codeartifact_listRepositoriesInDomainCmd.Flags().String("repository-prefix", "", "A prefix used to filter returned repositories.")
	codeartifact_listRepositoriesInDomainCmd.MarkFlagRequired("domain")
	codeartifactCmd.AddCommand(codeartifact_listRepositoriesInDomainCmd)
}
