package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var controlcatalog_listControlMappingsCmd = &cobra.Command{
	Use:   "list-control-mappings",
	Short: "Returns a paginated list of control mappings from the Control Catalog.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(controlcatalog_listControlMappingsCmd).Standalone()

	controlcatalog_listControlMappingsCmd.Flags().String("filter", "", "An optional filter that narrows the results to specific control mappings based on control ARNs, common control ARNs, or mapping types.")
	controlcatalog_listControlMappingsCmd.Flags().String("max-results", "", "The maximum number of results on a page or for an API request call.")
	controlcatalog_listControlMappingsCmd.Flags().String("next-token", "", "The pagination token that's used to fetch the next set of results.")
	controlcatalogCmd.AddCommand(controlcatalog_listControlMappingsCmd)
}
