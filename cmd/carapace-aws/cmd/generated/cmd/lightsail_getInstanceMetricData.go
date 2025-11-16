package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lightsail_getInstanceMetricDataCmd = &cobra.Command{
	Use:   "get-instance-metric-data",
	Short: "Returns the data points for the specified Amazon Lightsail instance metric, given an instance name.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lightsail_getInstanceMetricDataCmd).Standalone()

	lightsail_getInstanceMetricDataCmd.Flags().String("end-time", "", "The end time of the time period.")
	lightsail_getInstanceMetricDataCmd.Flags().String("instance-name", "", "The name of the instance for which you want to get metrics data.")
	lightsail_getInstanceMetricDataCmd.Flags().String("metric-name", "", "The metric for which you want to return information.")
	lightsail_getInstanceMetricDataCmd.Flags().String("period", "", "The granularity, in seconds, of the returned data points.")
	lightsail_getInstanceMetricDataCmd.Flags().String("start-time", "", "The start time of the time period.")
	lightsail_getInstanceMetricDataCmd.Flags().String("statistics", "", "The statistic for the metric.")
	lightsail_getInstanceMetricDataCmd.Flags().String("unit", "", "The unit for the metric data request.")
	lightsail_getInstanceMetricDataCmd.MarkFlagRequired("end-time")
	lightsail_getInstanceMetricDataCmd.MarkFlagRequired("instance-name")
	lightsail_getInstanceMetricDataCmd.MarkFlagRequired("metric-name")
	lightsail_getInstanceMetricDataCmd.MarkFlagRequired("period")
	lightsail_getInstanceMetricDataCmd.MarkFlagRequired("start-time")
	lightsail_getInstanceMetricDataCmd.MarkFlagRequired("statistics")
	lightsail_getInstanceMetricDataCmd.MarkFlagRequired("unit")
	lightsailCmd.AddCommand(lightsail_getInstanceMetricDataCmd)
}
