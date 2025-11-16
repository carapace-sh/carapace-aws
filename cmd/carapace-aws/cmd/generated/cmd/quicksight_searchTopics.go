package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_searchTopicsCmd = &cobra.Command{
	Use:   "search-topics",
	Short: "Searches for any Q topic that exists in an Quick Suite account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_searchTopicsCmd).Standalone()

	quicksight_searchTopicsCmd.Flags().String("aws-account-id", "", "The ID of the Amazon Web Services account that contains the topic that you want to find.")
	quicksight_searchTopicsCmd.Flags().String("filters", "", "The filters that you want to use to search for the topic.")
	quicksight_searchTopicsCmd.Flags().String("max-results", "", "The maximum number of results to be returned per request.")
	quicksight_searchTopicsCmd.Flags().String("next-token", "", "The token for the next set of results, or null if there are no more results.")
	quicksight_searchTopicsCmd.MarkFlagRequired("aws-account-id")
	quicksight_searchTopicsCmd.MarkFlagRequired("filters")
	quicksightCmd.AddCommand(quicksight_searchTopicsCmd)
}
