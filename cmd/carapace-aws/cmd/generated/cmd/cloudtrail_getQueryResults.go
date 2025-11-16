package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudtrail_getQueryResultsCmd = &cobra.Command{
	Use:   "get-query-results",
	Short: "Gets event data results of a query.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudtrail_getQueryResultsCmd).Standalone()

	cloudtrail_getQueryResultsCmd.Flags().String("event-data-store", "", "The ARN (or ID suffix of the ARN) of the event data store against which the query was run.")
	cloudtrail_getQueryResultsCmd.Flags().String("event-data-store-owner-account-id", "", "The account ID of the event data store owner.")
	cloudtrail_getQueryResultsCmd.Flags().String("max-query-results", "", "The maximum number of query results to display on a single page.")
	cloudtrail_getQueryResultsCmd.Flags().String("next-token", "", "A token you can use to get the next page of query results.")
	cloudtrail_getQueryResultsCmd.Flags().String("query-id", "", "The ID of the query for which you want to get results.")
	cloudtrail_getQueryResultsCmd.MarkFlagRequired("query-id")
	cloudtrailCmd.AddCommand(cloudtrail_getQueryResultsCmd)
}
