package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var internetmonitor_getQueryStatusCmd = &cobra.Command{
	Use:   "get-query-status",
	Short: "Returns the current status of a query for the Amazon CloudWatch Internet Monitor query interface, for a specified query ID and monitor.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(internetmonitor_getQueryStatusCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(internetmonitor_getQueryStatusCmd).Standalone()

		internetmonitor_getQueryStatusCmd.Flags().String("monitor-name", "", "The name of the monitor.")
		internetmonitor_getQueryStatusCmd.Flags().String("query-id", "", "The ID of the query that you want to return the status for.")
		internetmonitor_getQueryStatusCmd.MarkFlagRequired("monitor-name")
		internetmonitor_getQueryStatusCmd.MarkFlagRequired("query-id")
	})
	internetmonitorCmd.AddCommand(internetmonitor_getQueryStatusCmd)
}
