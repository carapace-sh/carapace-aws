package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rum_getAppMonitorDataCmd = &cobra.Command{
	Use:   "get-app-monitor-data",
	Short: "Retrieves the raw performance events that RUM has collected from your web application, so that you can do your own processing or analysis of this data.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rum_getAppMonitorDataCmd).Standalone()

	rum_getAppMonitorDataCmd.Flags().String("filters", "", "An array of structures that you can use to filter the results to those that match one or more sets of key-value pairs that you specify.")
	rum_getAppMonitorDataCmd.Flags().String("max-results", "", "The maximum number of results to return in one operation.")
	rum_getAppMonitorDataCmd.Flags().String("name", "", "The name of the app monitor that collected the data that you want to retrieve.")
	rum_getAppMonitorDataCmd.Flags().String("next-token", "", "Use the token returned by the previous operation to request the next page of results.")
	rum_getAppMonitorDataCmd.Flags().String("time-range", "", "A structure that defines the time range that you want to retrieve results from.")
	rum_getAppMonitorDataCmd.MarkFlagRequired("name")
	rum_getAppMonitorDataCmd.MarkFlagRequired("time-range")
	rumCmd.AddCommand(rum_getAppMonitorDataCmd)
}
