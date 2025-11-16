package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_getScalingConfigurationRecommendationCmd = &cobra.Command{
	Use:   "get-scaling-configuration-recommendation",
	Short: "Starts an Amazon SageMaker Inference Recommender autoscaling recommendation job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_getScalingConfigurationRecommendationCmd).Standalone()

	sagemaker_getScalingConfigurationRecommendationCmd.Flags().String("endpoint-name", "", "The name of an endpoint benchmarked during a previously completed inference recommendation job.")
	sagemaker_getScalingConfigurationRecommendationCmd.Flags().String("inference-recommendations-job-name", "", "The name of a previously completed Inference Recommender job.")
	sagemaker_getScalingConfigurationRecommendationCmd.Flags().String("recommendation-id", "", "The recommendation ID of a previously completed inference recommendation.")
	sagemaker_getScalingConfigurationRecommendationCmd.Flags().String("scaling-policy-objective", "", "An object where you specify the anticipated traffic pattern for an endpoint.")
	sagemaker_getScalingConfigurationRecommendationCmd.Flags().String("target-cpu-utilization-per-core", "", "The percentage of how much utilization you want an instance to use before autoscaling.")
	sagemaker_getScalingConfigurationRecommendationCmd.MarkFlagRequired("inference-recommendations-job-name")
	sagemakerCmd.AddCommand(sagemaker_getScalingConfigurationRecommendationCmd)
}
