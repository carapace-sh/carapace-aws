package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rum_listRumMetricsDestinationsCmd = &cobra.Command{
	Use:   "list-rum-metrics-destinations",
	Short: "Returns a list of destinations that you have created to receive RUM extended metrics, for the specified app monitor.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rum_listRumMetricsDestinationsCmd).Standalone()

	rum_listRumMetricsDestinationsCmd.Flags().String("app-monitor-name", "", "The name of the app monitor associated with the destinations that you want to retrieve.")
	rum_listRumMetricsDestinationsCmd.Flags().String("max-results", "", "The maximum number of results to return in one operation.")
	rum_listRumMetricsDestinationsCmd.Flags().String("next-token", "", "Use the token returned by the previous operation to request the next page of results.")
	rum_listRumMetricsDestinationsCmd.MarkFlagRequired("app-monitor-name")
	rumCmd.AddCommand(rum_listRumMetricsDestinationsCmd)
}
