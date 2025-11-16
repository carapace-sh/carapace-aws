package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_searchActionConnectorsCmd = &cobra.Command{
	Use:   "search-action-connectors",
	Short: "Searches for action connectors in the specified Amazon Web Services account using filters.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_searchActionConnectorsCmd).Standalone()

	quicksight_searchActionConnectorsCmd.Flags().String("aws-account-id", "", "The Amazon Web Services account ID in which to search for action connectors.")
	quicksight_searchActionConnectorsCmd.Flags().String("filters", "", "The search filters to apply.")
	quicksight_searchActionConnectorsCmd.Flags().String("max-results", "", "The maximum number of action connectors to return in a single response.")
	quicksight_searchActionConnectorsCmd.Flags().String("next-token", "", "A pagination token to retrieve the next set of results.")
	quicksight_searchActionConnectorsCmd.MarkFlagRequired("aws-account-id")
	quicksight_searchActionConnectorsCmd.MarkFlagRequired("filters")
	quicksightCmd.AddCommand(quicksight_searchActionConnectorsCmd)
}
