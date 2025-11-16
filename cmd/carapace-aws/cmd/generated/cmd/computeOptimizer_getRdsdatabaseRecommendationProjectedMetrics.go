package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var computeOptimizer_getRdsdatabaseRecommendationProjectedMetricsCmd = &cobra.Command{
	Use:   "get-rdsdatabase-recommendation-projected-metrics",
	Short: "Returns the projected metrics of Aurora and RDS database recommendations.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(computeOptimizer_getRdsdatabaseRecommendationProjectedMetricsCmd).Standalone()

	computeOptimizer_getRdsdatabaseRecommendationProjectedMetricsCmd.Flags().String("end-time", "", "The timestamp of the last projected metrics data point to return.")
	computeOptimizer_getRdsdatabaseRecommendationProjectedMetricsCmd.Flags().String("period", "", "The granularity, in seconds, of the projected metrics data points.")
	computeOptimizer_getRdsdatabaseRecommendationProjectedMetricsCmd.Flags().String("recommendation-preferences", "", "")
	computeOptimizer_getRdsdatabaseRecommendationProjectedMetricsCmd.Flags().String("resource-arn", "", "The ARN that identifies the Amazon Aurora or RDS database.")
	computeOptimizer_getRdsdatabaseRecommendationProjectedMetricsCmd.Flags().String("start-time", "", "The timestamp of the first projected metrics data point to return.")
	computeOptimizer_getRdsdatabaseRecommendationProjectedMetricsCmd.Flags().String("stat", "", "The statistic of the projected metrics.")
	computeOptimizer_getRdsdatabaseRecommendationProjectedMetricsCmd.MarkFlagRequired("end-time")
	computeOptimizer_getRdsdatabaseRecommendationProjectedMetricsCmd.MarkFlagRequired("period")
	computeOptimizer_getRdsdatabaseRecommendationProjectedMetricsCmd.MarkFlagRequired("resource-arn")
	computeOptimizer_getRdsdatabaseRecommendationProjectedMetricsCmd.MarkFlagRequired("start-time")
	computeOptimizer_getRdsdatabaseRecommendationProjectedMetricsCmd.MarkFlagRequired("stat")
	computeOptimizerCmd.AddCommand(computeOptimizer_getRdsdatabaseRecommendationProjectedMetricsCmd)
}
