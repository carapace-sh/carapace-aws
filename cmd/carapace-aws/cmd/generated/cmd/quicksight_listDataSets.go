package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_listDataSetsCmd = &cobra.Command{
	Use:   "list-data-sets",
	Short: "Lists all of the datasets belonging to the current Amazon Web Services account in an Amazon Web Services Region.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_listDataSetsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(quicksight_listDataSetsCmd).Standalone()

		quicksight_listDataSetsCmd.Flags().String("aws-account-id", "", "The Amazon Web Services account ID.")
		quicksight_listDataSetsCmd.Flags().String("max-results", "", "The maximum number of results to be returned per request.")
		quicksight_listDataSetsCmd.Flags().String("next-token", "", "The token for the next set of results, or null if there are no more results.")
		quicksight_listDataSetsCmd.MarkFlagRequired("aws-account-id")
	})
	quicksightCmd.AddCommand(quicksight_listDataSetsCmd)
}
