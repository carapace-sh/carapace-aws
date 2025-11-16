package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var applicationInsights_describeComponentConfigurationRecommendationCmd = &cobra.Command{
	Use:   "describe-component-configuration-recommendation",
	Short: "Describes the recommended monitoring configuration of the component.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(applicationInsights_describeComponentConfigurationRecommendationCmd).Standalone()

	applicationInsights_describeComponentConfigurationRecommendationCmd.Flags().String("component-name", "", "The name of the component.")
	applicationInsights_describeComponentConfigurationRecommendationCmd.Flags().String("recommendation-type", "", "The recommended configuration type.")
	applicationInsights_describeComponentConfigurationRecommendationCmd.Flags().String("resource-group-name", "", "The name of the resource group.")
	applicationInsights_describeComponentConfigurationRecommendationCmd.Flags().String("tier", "", "The tier of the application component.")
	applicationInsights_describeComponentConfigurationRecommendationCmd.Flags().String("workload-name", "", "The name of the workload.")
	applicationInsights_describeComponentConfigurationRecommendationCmd.MarkFlagRequired("component-name")
	applicationInsights_describeComponentConfigurationRecommendationCmd.MarkFlagRequired("resource-group-name")
	applicationInsights_describeComponentConfigurationRecommendationCmd.MarkFlagRequired("tier")
	applicationInsightsCmd.AddCommand(applicationInsights_describeComponentConfigurationRecommendationCmd)
}
