package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lakeformation_searchTablesByLftagsCmd = &cobra.Command{
	Use:   "search-tables-by-lftags",
	Short: "This operation allows a search on `TABLE` resources by `LFTag`s.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lakeformation_searchTablesByLftagsCmd).Standalone()

	lakeformation_searchTablesByLftagsCmd.Flags().String("catalog-id", "", "The identifier for the Data Catalog.")
	lakeformation_searchTablesByLftagsCmd.Flags().String("expression", "", "A list of conditions (`LFTag` structures) to search for in table resources.")
	lakeformation_searchTablesByLftagsCmd.Flags().String("max-results", "", "The maximum number of results to return.")
	lakeformation_searchTablesByLftagsCmd.Flags().String("next-token", "", "A continuation token, if this is not the first call to retrieve this list.")
	lakeformation_searchTablesByLftagsCmd.MarkFlagRequired("expression")
	lakeformationCmd.AddCommand(lakeformation_searchTablesByLftagsCmd)
}
