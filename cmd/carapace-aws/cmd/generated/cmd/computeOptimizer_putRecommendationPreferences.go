package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var computeOptimizer_putRecommendationPreferencesCmd = &cobra.Command{
	Use:   "put-recommendation-preferences",
	Short: "Creates a new recommendation preference or updates an existing recommendation preference, such as enhanced infrastructure metrics.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(computeOptimizer_putRecommendationPreferencesCmd).Standalone()

	computeOptimizer_putRecommendationPreferencesCmd.Flags().String("enhanced-infrastructure-metrics", "", "The status of the enhanced infrastructure metrics recommendation preference to create or update.")
	computeOptimizer_putRecommendationPreferencesCmd.Flags().String("external-metrics-preference", "", "The provider of the external metrics recommendation preference to create or update.")
	computeOptimizer_putRecommendationPreferencesCmd.Flags().String("inferred-workload-types", "", "The status of the inferred workload types recommendation preference to create or update.")
	computeOptimizer_putRecommendationPreferencesCmd.Flags().String("look-back-period", "", "The preference to control the number of days the utilization metrics of the Amazon Web Services resource are analyzed.")
	computeOptimizer_putRecommendationPreferencesCmd.Flags().String("preferred-resources", "", "The preference to control which resource type values are considered when generating rightsizing recommendations.")
	computeOptimizer_putRecommendationPreferencesCmd.Flags().String("resource-type", "", "The target resource type of the recommendation preference to create.")
	computeOptimizer_putRecommendationPreferencesCmd.Flags().String("savings-estimation-mode", "", "The status of the savings estimation mode preference to create or update.")
	computeOptimizer_putRecommendationPreferencesCmd.Flags().String("scope", "", "An object that describes the scope of the recommendation preference to create.")
	computeOptimizer_putRecommendationPreferencesCmd.Flags().String("utilization-preferences", "", "The preference to control the resource’s CPU utilization threshold, CPU utilization headroom, and memory utilization headroom.")
	computeOptimizer_putRecommendationPreferencesCmd.MarkFlagRequired("resource-type")
	computeOptimizerCmd.AddCommand(computeOptimizer_putRecommendationPreferencesCmd)
}
