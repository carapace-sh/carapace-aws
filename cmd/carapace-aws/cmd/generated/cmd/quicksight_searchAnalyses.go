package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_searchAnalysesCmd = &cobra.Command{
	Use:   "search-analyses",
	Short: "Searches for analyses that belong to the user specified in the filter.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_searchAnalysesCmd).Standalone()

	quicksight_searchAnalysesCmd.Flags().String("aws-account-id", "", "The ID of the Amazon Web Services account that contains the analyses that you're searching for.")
	quicksight_searchAnalysesCmd.Flags().String("filters", "", "The structure for the search filters that you want to apply to your search.")
	quicksight_searchAnalysesCmd.Flags().String("max-results", "", "The maximum number of results to return.")
	quicksight_searchAnalysesCmd.Flags().String("next-token", "", "A pagination token that can be used in a subsequent request.")
	quicksight_searchAnalysesCmd.MarkFlagRequired("aws-account-id")
	quicksight_searchAnalysesCmd.MarkFlagRequired("filters")
	quicksightCmd.AddCommand(quicksight_searchAnalysesCmd)
}
