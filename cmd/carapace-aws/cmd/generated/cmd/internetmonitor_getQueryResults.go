package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var internetmonitor_getQueryResultsCmd = &cobra.Command{
	Use:   "get-query-results",
	Short: "Return the data for a query with the Amazon CloudWatch Internet Monitor query interface.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(internetmonitor_getQueryResultsCmd).Standalone()

	internetmonitor_getQueryResultsCmd.Flags().String("max-results", "", "The number of query results that you want to return with this call.")
	internetmonitor_getQueryResultsCmd.Flags().String("monitor-name", "", "The name of the monitor to return data for.")
	internetmonitor_getQueryResultsCmd.Flags().String("next-token", "", "The token for the next set of results.")
	internetmonitor_getQueryResultsCmd.Flags().String("query-id", "", "The ID of the query that you want to return data results for.")
	internetmonitor_getQueryResultsCmd.MarkFlagRequired("monitor-name")
	internetmonitor_getQueryResultsCmd.MarkFlagRequired("query-id")
	internetmonitorCmd.AddCommand(internetmonitor_getQueryResultsCmd)
}
