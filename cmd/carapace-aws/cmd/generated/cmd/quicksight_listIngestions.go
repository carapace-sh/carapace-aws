package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_listIngestionsCmd = &cobra.Command{
	Use:   "list-ingestions",
	Short: "Lists the history of SPICE ingestions for a dataset.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_listIngestionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(quicksight_listIngestionsCmd).Standalone()

		quicksight_listIngestionsCmd.Flags().String("aws-account-id", "", "The Amazon Web Services account ID.")
		quicksight_listIngestionsCmd.Flags().String("data-set-id", "", "The ID of the dataset used in the ingestion.")
		quicksight_listIngestionsCmd.Flags().String("max-results", "", "The maximum number of results to be returned per request.")
		quicksight_listIngestionsCmd.Flags().String("next-token", "", "The token for the next set of results, or null if there are no more results.")
		quicksight_listIngestionsCmd.MarkFlagRequired("aws-account-id")
		quicksight_listIngestionsCmd.MarkFlagRequired("data-set-id")
	})
	quicksightCmd.AddCommand(quicksight_listIngestionsCmd)
}
