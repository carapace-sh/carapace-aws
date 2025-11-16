package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var costOptimizationHub_getRecommendationCmd = &cobra.Command{
	Use:   "get-recommendation",
	Short: "Returns both the current and recommended resource configuration and the estimated cost impact for a recommendation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(costOptimizationHub_getRecommendationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(costOptimizationHub_getRecommendationCmd).Standalone()

		costOptimizationHub_getRecommendationCmd.Flags().String("recommendation-id", "", "The ID for the recommendation.")
		costOptimizationHub_getRecommendationCmd.MarkFlagRequired("recommendation-id")
	})
	costOptimizationHubCmd.AddCommand(costOptimizationHub_getRecommendationCmd)
}
