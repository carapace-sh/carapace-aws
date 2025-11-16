package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_listDataSourcesCmd = &cobra.Command{
	Use:   "list-data-sources",
	Short: "Lists data sources in current Amazon Web Services Region that belong to this Amazon Web Services account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_listDataSourcesCmd).Standalone()

	quicksight_listDataSourcesCmd.Flags().String("aws-account-id", "", "The Amazon Web Services account ID.")
	quicksight_listDataSourcesCmd.Flags().String("max-results", "", "The maximum number of results to be returned per request.")
	quicksight_listDataSourcesCmd.Flags().String("next-token", "", "The token for the next set of results, or null if there are no more results.")
	quicksight_listDataSourcesCmd.MarkFlagRequired("aws-account-id")
	quicksightCmd.AddCommand(quicksight_listDataSourcesCmd)
}
