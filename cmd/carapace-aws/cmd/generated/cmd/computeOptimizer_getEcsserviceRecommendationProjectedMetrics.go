package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var computeOptimizer_getEcsserviceRecommendationProjectedMetricsCmd = &cobra.Command{
	Use:   "get-ecsservice-recommendation-projected-metrics",
	Short: "Returns the projected metrics of Amazon ECS service recommendations.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(computeOptimizer_getEcsserviceRecommendationProjectedMetricsCmd).Standalone()

	computeOptimizer_getEcsserviceRecommendationProjectedMetricsCmd.Flags().String("end-time", "", "The timestamp of the last projected metrics data point to return.")
	computeOptimizer_getEcsserviceRecommendationProjectedMetricsCmd.Flags().String("period", "", "The granularity, in seconds, of the projected metrics data points.")
	computeOptimizer_getEcsserviceRecommendationProjectedMetricsCmd.Flags().String("service-arn", "", "The ARN that identifies the Amazon ECS service.")
	computeOptimizer_getEcsserviceRecommendationProjectedMetricsCmd.Flags().String("start-time", "", "The timestamp of the first projected metrics data point to return.")
	computeOptimizer_getEcsserviceRecommendationProjectedMetricsCmd.Flags().String("stat", "", "The statistic of the projected metrics.")
	computeOptimizer_getEcsserviceRecommendationProjectedMetricsCmd.MarkFlagRequired("end-time")
	computeOptimizer_getEcsserviceRecommendationProjectedMetricsCmd.MarkFlagRequired("period")
	computeOptimizer_getEcsserviceRecommendationProjectedMetricsCmd.MarkFlagRequired("service-arn")
	computeOptimizer_getEcsserviceRecommendationProjectedMetricsCmd.MarkFlagRequired("start-time")
	computeOptimizer_getEcsserviceRecommendationProjectedMetricsCmd.MarkFlagRequired("stat")
	computeOptimizerCmd.AddCommand(computeOptimizer_getEcsserviceRecommendationProjectedMetricsCmd)
}
