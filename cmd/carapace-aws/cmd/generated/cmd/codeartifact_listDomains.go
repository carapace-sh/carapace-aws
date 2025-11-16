package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codeartifact_listDomainsCmd = &cobra.Command{
	Use:   "list-domains",
	Short: "Returns a list of [DomainSummary](https://docs.aws.amazon.com/codeartifact/latest/APIReference/API_PackageVersionDescription.html) objects for all domains owned by the Amazon Web Services account that makes this call.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codeartifact_listDomainsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codeartifact_listDomainsCmd).Standalone()

		codeartifact_listDomainsCmd.Flags().String("max-results", "", "The maximum number of results to return per page.")
		codeartifact_listDomainsCmd.Flags().String("next-token", "", "The token for the next set of results.")
	})
	codeartifactCmd.AddCommand(codeartifact_listDomainsCmd)
}
