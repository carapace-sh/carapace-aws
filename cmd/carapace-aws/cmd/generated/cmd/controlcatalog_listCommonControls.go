package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var controlcatalog_listCommonControlsCmd = &cobra.Command{
	Use:   "list-common-controls",
	Short: "Returns a paginated list of common controls from the Amazon Web Services Control Catalog.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(controlcatalog_listCommonControlsCmd).Standalone()

	controlcatalog_listCommonControlsCmd.Flags().String("common-control-filter", "", "An optional filter that narrows the results to a specific objective.")
	controlcatalog_listCommonControlsCmd.Flags().String("max-results", "", "The maximum number of results on a page or for an API request call.")
	controlcatalog_listCommonControlsCmd.Flags().String("next-token", "", "The pagination token that's used to fetch the next set of results.")
	controlcatalogCmd.AddCommand(controlcatalog_listCommonControlsCmd)
}
