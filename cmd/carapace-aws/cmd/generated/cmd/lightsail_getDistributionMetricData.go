package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lightsail_getDistributionMetricDataCmd = &cobra.Command{
	Use:   "get-distribution-metric-data",
	Short: "Returns the data points of a specific metric for an Amazon Lightsail content delivery network (CDN) distribution.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lightsail_getDistributionMetricDataCmd).Standalone()

	lightsail_getDistributionMetricDataCmd.Flags().String("distribution-name", "", "The name of the distribution for which to get metric data.")
	lightsail_getDistributionMetricDataCmd.Flags().String("end-time", "", "The end of the time interval for which to get metric data.")
	lightsail_getDistributionMetricDataCmd.Flags().String("metric-name", "", "The metric for which you want to return information.")
	lightsail_getDistributionMetricDataCmd.Flags().String("period", "", "The granularity, in seconds, for the metric data points that will be returned.")
	lightsail_getDistributionMetricDataCmd.Flags().String("start-time", "", "The start of the time interval for which to get metric data.")
	lightsail_getDistributionMetricDataCmd.Flags().String("statistics", "", "The statistic for the metric.")
	lightsail_getDistributionMetricDataCmd.Flags().String("unit", "", "The unit for the metric data request.")
	lightsail_getDistributionMetricDataCmd.MarkFlagRequired("distribution-name")
	lightsail_getDistributionMetricDataCmd.MarkFlagRequired("end-time")
	lightsail_getDistributionMetricDataCmd.MarkFlagRequired("metric-name")
	lightsail_getDistributionMetricDataCmd.MarkFlagRequired("period")
	lightsail_getDistributionMetricDataCmd.MarkFlagRequired("start-time")
	lightsail_getDistributionMetricDataCmd.MarkFlagRequired("statistics")
	lightsail_getDistributionMetricDataCmd.MarkFlagRequired("unit")
	lightsailCmd.AddCommand(lightsail_getDistributionMetricDataCmd)
}
