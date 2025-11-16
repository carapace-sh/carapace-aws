package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_searchDataSourcesCmd = &cobra.Command{
	Use:   "search-data-sources",
	Short: "Use the `SearchDataSources` operation to search for data sources that belong to an account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_searchDataSourcesCmd).Standalone()

	quicksight_searchDataSourcesCmd.Flags().String("aws-account-id", "", "The Amazon Web Services account ID.")
	quicksight_searchDataSourcesCmd.Flags().String("filters", "", "The filters to apply to the search.")
	quicksight_searchDataSourcesCmd.Flags().String("max-results", "", "The maximum number of results to be returned per request.")
	quicksight_searchDataSourcesCmd.Flags().String("next-token", "", "A pagination token that can be used in a subsequent request.")
	quicksight_searchDataSourcesCmd.MarkFlagRequired("aws-account-id")
	quicksight_searchDataSourcesCmd.MarkFlagRequired("filters")
	quicksightCmd.AddCommand(quicksight_searchDataSourcesCmd)
}
