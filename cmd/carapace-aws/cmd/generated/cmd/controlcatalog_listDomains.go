package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var controlcatalog_listDomainsCmd = &cobra.Command{
	Use:   "list-domains",
	Short: "Returns a paginated list of domains from the Control Catalog.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(controlcatalog_listDomainsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(controlcatalog_listDomainsCmd).Standalone()

		controlcatalog_listDomainsCmd.Flags().String("max-results", "", "The maximum number of results on a page or for an API request call.")
		controlcatalog_listDomainsCmd.Flags().String("next-token", "", "The pagination token that's used to fetch the next set of results.")
	})
	controlcatalogCmd.AddCommand(controlcatalog_listDomainsCmd)
}
