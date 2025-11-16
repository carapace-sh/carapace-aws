package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var computeOptimizer_getEc2RecommendationProjectedMetricsCmd = &cobra.Command{
	Use:   "get-ec2-recommendation-projected-metrics",
	Short: "Returns the projected utilization metrics of Amazon EC2 instance recommendations.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(computeOptimizer_getEc2RecommendationProjectedMetricsCmd).Standalone()

	computeOptimizer_getEc2RecommendationProjectedMetricsCmd.Flags().String("end-time", "", "The timestamp of the last projected metrics data point to return.")
	computeOptimizer_getEc2RecommendationProjectedMetricsCmd.Flags().String("instance-arn", "", "The Amazon Resource Name (ARN) of the instances for which to return recommendation projected metrics.")
	computeOptimizer_getEc2RecommendationProjectedMetricsCmd.Flags().String("period", "", "The granularity, in seconds, of the projected metrics data points.")
	computeOptimizer_getEc2RecommendationProjectedMetricsCmd.Flags().String("recommendation-preferences", "", "An object to specify the preferences for the Amazon EC2 recommendation projected metrics to return in the response.")
	computeOptimizer_getEc2RecommendationProjectedMetricsCmd.Flags().String("start-time", "", "The timestamp of the first projected metrics data point to return.")
	computeOptimizer_getEc2RecommendationProjectedMetricsCmd.Flags().String("stat", "", "The statistic of the projected metrics.")
	computeOptimizer_getEc2RecommendationProjectedMetricsCmd.MarkFlagRequired("end-time")
	computeOptimizer_getEc2RecommendationProjectedMetricsCmd.MarkFlagRequired("instance-arn")
	computeOptimizer_getEc2RecommendationProjectedMetricsCmd.MarkFlagRequired("period")
	computeOptimizer_getEc2RecommendationProjectedMetricsCmd.MarkFlagRequired("start-time")
	computeOptimizer_getEc2RecommendationProjectedMetricsCmd.MarkFlagRequired("stat")
	computeOptimizerCmd.AddCommand(computeOptimizer_getEc2RecommendationProjectedMetricsCmd)
}
