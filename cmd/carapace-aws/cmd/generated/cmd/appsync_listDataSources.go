package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appsync_listDataSourcesCmd = &cobra.Command{
	Use:   "list-data-sources",
	Short: "Lists the data sources for a given API.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appsync_listDataSourcesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(appsync_listDataSourcesCmd).Standalone()

		appsync_listDataSourcesCmd.Flags().String("api-id", "", "The API ID.")
		appsync_listDataSourcesCmd.Flags().String("max-results", "", "The maximum number of results that you want the request to return.")
		appsync_listDataSourcesCmd.Flags().String("next-token", "", "An identifier that was returned from the previous call to this operation, which you can use to return the next set of items in the list.")
		appsync_listDataSourcesCmd.MarkFlagRequired("api-id")
	})
	appsyncCmd.AddCommand(appsync_listDataSourcesCmd)
}
