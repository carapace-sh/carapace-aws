package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudwatch_getMetricStatisticsCmd = &cobra.Command{
	Use:   "get-metric-statistics",
	Short: "Gets statistics for the specified metric.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudwatch_getMetricStatisticsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloudwatch_getMetricStatisticsCmd).Standalone()

		cloudwatch_getMetricStatisticsCmd.Flags().String("dimensions", "", "The dimensions.")
		cloudwatch_getMetricStatisticsCmd.Flags().String("end-time", "", "The time stamp that determines the last data point to return.")
		cloudwatch_getMetricStatisticsCmd.Flags().String("extended-statistics", "", "The percentile statistics.")
		cloudwatch_getMetricStatisticsCmd.Flags().String("metric-name", "", "The name of the metric, with or without spaces.")
		cloudwatch_getMetricStatisticsCmd.Flags().String("namespace", "", "The namespace of the metric, with or without spaces.")
		cloudwatch_getMetricStatisticsCmd.Flags().String("period", "", "The granularity, in seconds, of the returned data points.")
		cloudwatch_getMetricStatisticsCmd.Flags().String("start-time", "", "The time stamp that determines the first data point to return.")
		cloudwatch_getMetricStatisticsCmd.Flags().String("statistics", "", "The metric statistics, other than percentile.")
		cloudwatch_getMetricStatisticsCmd.Flags().String("unit", "", "The unit for a given metric.")
		cloudwatch_getMetricStatisticsCmd.MarkFlagRequired("end-time")
		cloudwatch_getMetricStatisticsCmd.MarkFlagRequired("metric-name")
		cloudwatch_getMetricStatisticsCmd.MarkFlagRequired("namespace")
		cloudwatch_getMetricStatisticsCmd.MarkFlagRequired("period")
		cloudwatch_getMetricStatisticsCmd.MarkFlagRequired("start-time")
	})
	cloudwatchCmd.AddCommand(cloudwatch_getMetricStatisticsCmd)
}
