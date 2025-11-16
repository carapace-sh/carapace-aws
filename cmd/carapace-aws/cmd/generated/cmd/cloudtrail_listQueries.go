package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudtrail_listQueriesCmd = &cobra.Command{
	Use:   "list-queries",
	Short: "Returns a list of queries and query statuses for the past seven days.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudtrail_listQueriesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloudtrail_listQueriesCmd).Standalone()

		cloudtrail_listQueriesCmd.Flags().String("end-time", "", "Use with `StartTime` to bound a `ListQueries` request, and limit its results to only those queries run within a specified time period.")
		cloudtrail_listQueriesCmd.Flags().String("event-data-store", "", "The ARN (or the ID suffix of the ARN) of an event data store on which queries were run.")
		cloudtrail_listQueriesCmd.Flags().String("max-results", "", "The maximum number of queries to show on a page.")
		cloudtrail_listQueriesCmd.Flags().String("next-token", "", "A token you can use to get the next page of results.")
		cloudtrail_listQueriesCmd.Flags().String("query-status", "", "The status of queries that you want to return in results.")
		cloudtrail_listQueriesCmd.Flags().String("start-time", "", "Use with `EndTime` to bound a `ListQueries` request, and limit its results to only those queries run within a specified time period.")
		cloudtrail_listQueriesCmd.MarkFlagRequired("event-data-store")
	})
	cloudtrailCmd.AddCommand(cloudtrail_listQueriesCmd)
}
