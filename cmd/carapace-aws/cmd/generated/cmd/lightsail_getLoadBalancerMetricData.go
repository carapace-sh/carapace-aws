package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lightsail_getLoadBalancerMetricDataCmd = &cobra.Command{
	Use:   "get-load-balancer-metric-data",
	Short: "Returns information about health metrics for your Lightsail load balancer.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lightsail_getLoadBalancerMetricDataCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lightsail_getLoadBalancerMetricDataCmd).Standalone()

		lightsail_getLoadBalancerMetricDataCmd.Flags().String("end-time", "", "The end time of the period.")
		lightsail_getLoadBalancerMetricDataCmd.Flags().String("load-balancer-name", "", "The name of the load balancer.")
		lightsail_getLoadBalancerMetricDataCmd.Flags().String("metric-name", "", "The metric for which you want to return information.")
		lightsail_getLoadBalancerMetricDataCmd.Flags().String("period", "", "The granularity, in seconds, of the returned data points.")
		lightsail_getLoadBalancerMetricDataCmd.Flags().String("start-time", "", "The start time of the period.")
		lightsail_getLoadBalancerMetricDataCmd.Flags().String("statistics", "", "The statistic for the metric.")
		lightsail_getLoadBalancerMetricDataCmd.Flags().String("unit", "", "The unit for the metric data request.")
		lightsail_getLoadBalancerMetricDataCmd.MarkFlagRequired("end-time")
		lightsail_getLoadBalancerMetricDataCmd.MarkFlagRequired("load-balancer-name")
		lightsail_getLoadBalancerMetricDataCmd.MarkFlagRequired("metric-name")
		lightsail_getLoadBalancerMetricDataCmd.MarkFlagRequired("period")
		lightsail_getLoadBalancerMetricDataCmd.MarkFlagRequired("start-time")
		lightsail_getLoadBalancerMetricDataCmd.MarkFlagRequired("statistics")
		lightsail_getLoadBalancerMetricDataCmd.MarkFlagRequired("unit")
	})
	lightsailCmd.AddCommand(lightsail_getLoadBalancerMetricDataCmd)
}
