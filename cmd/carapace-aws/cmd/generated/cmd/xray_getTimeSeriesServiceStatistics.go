package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var xray_getTimeSeriesServiceStatisticsCmd = &cobra.Command{
	Use:   "get-time-series-service-statistics",
	Short: "Get an aggregation of service statistics defined by a specific time range.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(xray_getTimeSeriesServiceStatisticsCmd).Standalone()

	xray_getTimeSeriesServiceStatisticsCmd.Flags().String("end-time", "", "The end of the time frame for which to aggregate statistics.")
	xray_getTimeSeriesServiceStatisticsCmd.Flags().String("entity-selector-expression", "", "A filter expression defining entities that will be aggregated for statistics.")
	xray_getTimeSeriesServiceStatisticsCmd.Flags().String("forecast-statistics", "", "The forecasted high and low fault count values.")
	xray_getTimeSeriesServiceStatisticsCmd.Flags().String("group-arn", "", "The Amazon Resource Name (ARN) of the group for which to pull statistics from.")
	xray_getTimeSeriesServiceStatisticsCmd.Flags().String("group-name", "", "The case-sensitive name of the group for which to pull statistics from.")
	xray_getTimeSeriesServiceStatisticsCmd.Flags().String("next-token", "", "Pagination token.")
	xray_getTimeSeriesServiceStatisticsCmd.Flags().String("period", "", "Aggregation period in seconds.")
	xray_getTimeSeriesServiceStatisticsCmd.Flags().String("start-time", "", "The start of the time frame for which to aggregate statistics.")
	xray_getTimeSeriesServiceStatisticsCmd.MarkFlagRequired("end-time")
	xray_getTimeSeriesServiceStatisticsCmd.MarkFlagRequired("start-time")
	xrayCmd.AddCommand(xray_getTimeSeriesServiceStatisticsCmd)
}
