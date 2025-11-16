package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_searchFlowsCmd = &cobra.Command{
	Use:   "search-flows",
	Short: "Search for the flows in an Amazon Web Services account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_searchFlowsCmd).Standalone()

	quicksight_searchFlowsCmd.Flags().String("aws-account-id", "", "The ID of the Amazon Web Services account where you are searching for flows from.")
	quicksight_searchFlowsCmd.Flags().String("filters", "", "The filters applied to the search when searching for flows in the Amazon Web Services account.")
	quicksight_searchFlowsCmd.Flags().String("max-results", "", "The maximum number of results to be returned per request.")
	quicksight_searchFlowsCmd.Flags().String("next-token", "", "The token to request the next set of results, or null if you want to retrieve the first set.")
	quicksight_searchFlowsCmd.MarkFlagRequired("aws-account-id")
	quicksight_searchFlowsCmd.MarkFlagRequired("filters")
	quicksightCmd.AddCommand(quicksight_searchFlowsCmd)
}
