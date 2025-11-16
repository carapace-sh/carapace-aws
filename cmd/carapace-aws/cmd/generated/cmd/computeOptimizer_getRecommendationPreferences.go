package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var computeOptimizer_getRecommendationPreferencesCmd = &cobra.Command{
	Use:   "get-recommendation-preferences",
	Short: "Returns existing recommendation preferences, such as enhanced infrastructure metrics.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(computeOptimizer_getRecommendationPreferencesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(computeOptimizer_getRecommendationPreferencesCmd).Standalone()

		computeOptimizer_getRecommendationPreferencesCmd.Flags().String("max-results", "", "The maximum number of recommendation preferences to return with a single request.")
		computeOptimizer_getRecommendationPreferencesCmd.Flags().String("next-token", "", "The token to advance to the next page of recommendation preferences.")
		computeOptimizer_getRecommendationPreferencesCmd.Flags().String("resource-type", "", "The target resource type of the recommendation preference for which to return preferences.")
		computeOptimizer_getRecommendationPreferencesCmd.Flags().String("scope", "", "An object that describes the scope of the recommendation preference to return.")
		computeOptimizer_getRecommendationPreferencesCmd.MarkFlagRequired("resource-type")
	})
	computeOptimizerCmd.AddCommand(computeOptimizer_getRecommendationPreferencesCmd)
}
