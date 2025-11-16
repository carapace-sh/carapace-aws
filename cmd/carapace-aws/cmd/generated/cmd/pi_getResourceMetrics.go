package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pi_getResourceMetricsCmd = &cobra.Command{
	Use:   "get-resource-metrics",
	Short: "Retrieve Performance Insights metrics for a set of data sources over a time period.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pi_getResourceMetricsCmd).Standalone()

	pi_getResourceMetricsCmd.Flags().String("end-time", "", "The date and time specifying the end of the requested time series query range.")
	pi_getResourceMetricsCmd.Flags().String("identifier", "", "An immutable identifier for a data source that is unique for an Amazon Web Services Region.")
	pi_getResourceMetricsCmd.Flags().String("max-results", "", "The maximum number of items to return in the response.")
	pi_getResourceMetricsCmd.Flags().String("metric-queries", "", "An array of one or more queries to perform.")
	pi_getResourceMetricsCmd.Flags().String("next-token", "", "An optional pagination token provided by a previous request.")
	pi_getResourceMetricsCmd.Flags().String("period-alignment", "", "The returned timestamp which is the start or end time of the time periods.")
	pi_getResourceMetricsCmd.Flags().String("period-in-seconds", "", "The granularity, in seconds, of the data points returned from Performance Insights.")
	pi_getResourceMetricsCmd.Flags().String("service-type", "", "The Amazon Web Services service for which Performance Insights returns metrics.")
	pi_getResourceMetricsCmd.Flags().String("start-time", "", "The date and time specifying the beginning of the requested time series query range.")
	pi_getResourceMetricsCmd.MarkFlagRequired("end-time")
	pi_getResourceMetricsCmd.MarkFlagRequired("identifier")
	pi_getResourceMetricsCmd.MarkFlagRequired("metric-queries")
	pi_getResourceMetricsCmd.MarkFlagRequired("service-type")
	pi_getResourceMetricsCmd.MarkFlagRequired("start-time")
	piCmd.AddCommand(pi_getResourceMetricsCmd)
}
