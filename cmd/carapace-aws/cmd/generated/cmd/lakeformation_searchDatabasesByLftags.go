package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lakeformation_searchDatabasesByLftagsCmd = &cobra.Command{
	Use:   "search-databases-by-lftags",
	Short: "This operation allows a search on `DATABASE` resources by `TagCondition`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lakeformation_searchDatabasesByLftagsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lakeformation_searchDatabasesByLftagsCmd).Standalone()

		lakeformation_searchDatabasesByLftagsCmd.Flags().String("catalog-id", "", "The identifier for the Data Catalog.")
		lakeformation_searchDatabasesByLftagsCmd.Flags().String("expression", "", "A list of conditions (`LFTag` structures) to search for in database resources.")
		lakeformation_searchDatabasesByLftagsCmd.Flags().String("max-results", "", "The maximum number of results to return.")
		lakeformation_searchDatabasesByLftagsCmd.Flags().String("next-token", "", "A continuation token, if this is not the first call to retrieve this list.")
		lakeformation_searchDatabasesByLftagsCmd.MarkFlagRequired("expression")
	})
	lakeformationCmd.AddCommand(lakeformation_searchDatabasesByLftagsCmd)
}
