package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var computeOptimizer_getEffectiveRecommendationPreferencesCmd = &cobra.Command{
	Use:   "get-effective-recommendation-preferences",
	Short: "Returns the recommendation preferences that are in effect for a given resource, such as enhanced infrastructure metrics.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(computeOptimizer_getEffectiveRecommendationPreferencesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(computeOptimizer_getEffectiveRecommendationPreferencesCmd).Standalone()

		computeOptimizer_getEffectiveRecommendationPreferencesCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource for which to confirm effective recommendation preferences.")
		computeOptimizer_getEffectiveRecommendationPreferencesCmd.MarkFlagRequired("resource-arn")
	})
	computeOptimizerCmd.AddCommand(computeOptimizer_getEffectiveRecommendationPreferencesCmd)
}
