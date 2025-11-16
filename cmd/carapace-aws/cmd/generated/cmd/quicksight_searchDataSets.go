package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_searchDataSetsCmd = &cobra.Command{
	Use:   "search-data-sets",
	Short: "Use the `SearchDataSets` operation to search for datasets that belong to an account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_searchDataSetsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(quicksight_searchDataSetsCmd).Standalone()

		quicksight_searchDataSetsCmd.Flags().String("aws-account-id", "", "The Amazon Web Services account ID.")
		quicksight_searchDataSetsCmd.Flags().String("filters", "", "The filters to apply to the search.")
		quicksight_searchDataSetsCmd.Flags().String("max-results", "", "The maximum number of results to be returned per request.")
		quicksight_searchDataSetsCmd.Flags().String("next-token", "", "A pagination token that can be used in a subsequent request.")
		quicksight_searchDataSetsCmd.MarkFlagRequired("aws-account-id")
		quicksight_searchDataSetsCmd.MarkFlagRequired("filters")
	})
	quicksightCmd.AddCommand(quicksight_searchDataSetsCmd)
}
