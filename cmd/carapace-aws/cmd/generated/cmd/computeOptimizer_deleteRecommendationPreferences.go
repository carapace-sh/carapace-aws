package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var computeOptimizer_deleteRecommendationPreferencesCmd = &cobra.Command{
	Use:   "delete-recommendation-preferences",
	Short: "Deletes a recommendation preference, such as enhanced infrastructure metrics.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(computeOptimizer_deleteRecommendationPreferencesCmd).Standalone()

	computeOptimizer_deleteRecommendationPreferencesCmd.Flags().String("recommendation-preference-names", "", "The name of the recommendation preference to delete.")
	computeOptimizer_deleteRecommendationPreferencesCmd.Flags().String("resource-type", "", "The target resource type of the recommendation preference to delete.")
	computeOptimizer_deleteRecommendationPreferencesCmd.Flags().String("scope", "", "An object that describes the scope of the recommendation preference to delete.")
	computeOptimizer_deleteRecommendationPreferencesCmd.MarkFlagRequired("recommendation-preference-names")
	computeOptimizer_deleteRecommendationPreferencesCmd.MarkFlagRequired("resource-type")
	computeOptimizerCmd.AddCommand(computeOptimizer_deleteRecommendationPreferencesCmd)
}
