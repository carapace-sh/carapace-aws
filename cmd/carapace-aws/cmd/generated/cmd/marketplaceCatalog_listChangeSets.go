package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var marketplaceCatalog_listChangeSetsCmd = &cobra.Command{
	Use:   "list-change-sets",
	Short: "Returns the list of change sets owned by the account being used to make the call.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(marketplaceCatalog_listChangeSetsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(marketplaceCatalog_listChangeSetsCmd).Standalone()

		marketplaceCatalog_listChangeSetsCmd.Flags().String("catalog", "", "The catalog related to the request.")
		marketplaceCatalog_listChangeSetsCmd.Flags().String("filter-list", "", "An array of filter objects.")
		marketplaceCatalog_listChangeSetsCmd.Flags().String("max-results", "", "The maximum number of results returned by a single call.")
		marketplaceCatalog_listChangeSetsCmd.Flags().String("next-token", "", "The token value retrieved from a previous call to access the next page of results.")
		marketplaceCatalog_listChangeSetsCmd.Flags().String("sort", "", "An object that contains two attributes, `SortBy` and `SortOrder`.")
		marketplaceCatalog_listChangeSetsCmd.MarkFlagRequired("catalog")
	})
	marketplaceCatalogCmd.AddCommand(marketplaceCatalog_listChangeSetsCmd)
}
