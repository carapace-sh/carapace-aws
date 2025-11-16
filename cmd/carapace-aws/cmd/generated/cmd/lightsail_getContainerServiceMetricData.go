package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lightsail_getContainerServiceMetricDataCmd = &cobra.Command{
	Use:   "get-container-service-metric-data",
	Short: "Returns the data points of a specific metric of your Amazon Lightsail container service.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lightsail_getContainerServiceMetricDataCmd).Standalone()

	lightsail_getContainerServiceMetricDataCmd.Flags().String("end-time", "", "The end time of the time period.")
	lightsail_getContainerServiceMetricDataCmd.Flags().String("metric-name", "", "The metric for which you want to return information.")
	lightsail_getContainerServiceMetricDataCmd.Flags().String("period", "", "The granularity, in seconds, of the returned data points.")
	lightsail_getContainerServiceMetricDataCmd.Flags().String("service-name", "", "The name of the container service for which to get metric data.")
	lightsail_getContainerServiceMetricDataCmd.Flags().String("start-time", "", "The start time of the time period.")
	lightsail_getContainerServiceMetricDataCmd.Flags().String("statistics", "", "The statistic for the metric.")
	lightsail_getContainerServiceMetricDataCmd.MarkFlagRequired("end-time")
	lightsail_getContainerServiceMetricDataCmd.MarkFlagRequired("metric-name")
	lightsail_getContainerServiceMetricDataCmd.MarkFlagRequired("period")
	lightsail_getContainerServiceMetricDataCmd.MarkFlagRequired("service-name")
	lightsail_getContainerServiceMetricDataCmd.MarkFlagRequired("start-time")
	lightsail_getContainerServiceMetricDataCmd.MarkFlagRequired("statistics")
	lightsailCmd.AddCommand(lightsail_getContainerServiceMetricDataCmd)
}
