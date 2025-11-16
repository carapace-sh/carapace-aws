package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var controlcatalog_listObjectivesCmd = &cobra.Command{
	Use:   "list-objectives",
	Short: "Returns a paginated list of objectives from the Control Catalog.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(controlcatalog_listObjectivesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(controlcatalog_listObjectivesCmd).Standalone()

		controlcatalog_listObjectivesCmd.Flags().String("max-results", "", "The maximum number of results on a page or for an API request call.")
		controlcatalog_listObjectivesCmd.Flags().String("next-token", "", "The pagination token that's used to fetch the next set of results.")
		controlcatalog_listObjectivesCmd.Flags().String("objective-filter", "", "An optional filter that narrows the results to a specific domain.")
	})
	controlcatalogCmd.AddCommand(controlcatalog_listObjectivesCmd)
}
