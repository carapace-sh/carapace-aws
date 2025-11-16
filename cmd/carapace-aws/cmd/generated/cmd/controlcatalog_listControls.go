package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var controlcatalog_listControlsCmd = &cobra.Command{
	Use:   "list-controls",
	Short: "Returns a paginated list of all available controls in the Control Catalog library.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(controlcatalog_listControlsCmd).Standalone()

	controlcatalog_listControlsCmd.Flags().String("filter", "", "An optional filter that narrows the results to controls with specific implementation types or identifiers.")
	controlcatalog_listControlsCmd.Flags().String("max-results", "", "The maximum number of results on a page or for an API request call.")
	controlcatalog_listControlsCmd.Flags().String("next-token", "", "The pagination token that's used to fetch the next set of results.")
	controlcatalogCmd.AddCommand(controlcatalog_listControlsCmd)
}
