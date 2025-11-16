package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var internetmonitor_startQueryCmd = &cobra.Command{
	Use:   "start-query",
	Short: "Start a query to return data for a specific query type for the Amazon CloudWatch Internet Monitor query interface.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(internetmonitor_startQueryCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(internetmonitor_startQueryCmd).Standalone()

		internetmonitor_startQueryCmd.Flags().String("end-time", "", "The timestamp that is the end of the period that you want to retrieve data for with your query.")
		internetmonitor_startQueryCmd.Flags().String("filter-parameters", "", "The `FilterParameters` field that you use with Amazon CloudWatch Internet Monitor queries is a string the defines how you want a query to be filtered.")
		internetmonitor_startQueryCmd.Flags().String("linked-account-id", "", "The account ID for an account that you've set up cross-account sharing for in Amazon CloudWatch Internet Monitor.")
		internetmonitor_startQueryCmd.Flags().String("monitor-name", "", "The name of the monitor to query.")
		internetmonitor_startQueryCmd.Flags().String("query-type", "", "The type of query to run.")
		internetmonitor_startQueryCmd.Flags().String("start-time", "", "The timestamp that is the beginning of the period that you want to retrieve data for with your query.")
		internetmonitor_startQueryCmd.MarkFlagRequired("end-time")
		internetmonitor_startQueryCmd.MarkFlagRequired("monitor-name")
		internetmonitor_startQueryCmd.MarkFlagRequired("query-type")
		internetmonitor_startQueryCmd.MarkFlagRequired("start-time")
	})
	internetmonitorCmd.AddCommand(internetmonitor_startQueryCmd)
}
