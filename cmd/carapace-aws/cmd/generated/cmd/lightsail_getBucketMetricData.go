package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lightsail_getBucketMetricDataCmd = &cobra.Command{
	Use:   "get-bucket-metric-data",
	Short: "Returns the data points of a specific metric for an Amazon Lightsail bucket.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lightsail_getBucketMetricDataCmd).Standalone()

	lightsail_getBucketMetricDataCmd.Flags().String("bucket-name", "", "The name of the bucket for which to get metric data.")
	lightsail_getBucketMetricDataCmd.Flags().String("end-time", "", "The timestamp indicating the latest data to be returned.")
	lightsail_getBucketMetricDataCmd.Flags().String("metric-name", "", "The metric for which you want to return information.")
	lightsail_getBucketMetricDataCmd.Flags().String("period", "", "The granularity, in seconds, of the returned data points.")
	lightsail_getBucketMetricDataCmd.Flags().String("start-time", "", "The timestamp indicating the earliest data to be returned.")
	lightsail_getBucketMetricDataCmd.Flags().String("statistics", "", "The statistic for the metric.")
	lightsail_getBucketMetricDataCmd.Flags().String("unit", "", "The unit for the metric data request.")
	lightsail_getBucketMetricDataCmd.MarkFlagRequired("bucket-name")
	lightsail_getBucketMetricDataCmd.MarkFlagRequired("end-time")
	lightsail_getBucketMetricDataCmd.MarkFlagRequired("metric-name")
	lightsail_getBucketMetricDataCmd.MarkFlagRequired("period")
	lightsail_getBucketMetricDataCmd.MarkFlagRequired("start-time")
	lightsail_getBucketMetricDataCmd.MarkFlagRequired("statistics")
	lightsail_getBucketMetricDataCmd.MarkFlagRequired("unit")
	lightsailCmd.AddCommand(lightsail_getBucketMetricDataCmd)
}
