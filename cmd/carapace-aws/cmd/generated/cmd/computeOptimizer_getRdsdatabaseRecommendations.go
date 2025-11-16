package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var computeOptimizer_getRdsdatabaseRecommendationsCmd = &cobra.Command{
	Use:   "get-rdsdatabase-recommendations",
	Short: "Returns Amazon Aurora and RDS database recommendations.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(computeOptimizer_getRdsdatabaseRecommendationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(computeOptimizer_getRdsdatabaseRecommendationsCmd).Standalone()

		computeOptimizer_getRdsdatabaseRecommendationsCmd.Flags().String("account-ids", "", "Return the Amazon Aurora and RDS database recommendations to the specified Amazon Web Services account IDs.")
		computeOptimizer_getRdsdatabaseRecommendationsCmd.Flags().String("filters", "", "An array of objects to specify a filter that returns a more specific list of Amazon Aurora and RDS database recommendations.")
		computeOptimizer_getRdsdatabaseRecommendationsCmd.Flags().String("max-results", "", "The maximum number of Amazon Aurora and RDS database recommendations to return with a single request.")
		computeOptimizer_getRdsdatabaseRecommendationsCmd.Flags().String("next-token", "", "The token to advance to the next page of Amazon Aurora and RDS database recommendations.")
		computeOptimizer_getRdsdatabaseRecommendationsCmd.Flags().String("recommendation-preferences", "", "")
		computeOptimizer_getRdsdatabaseRecommendationsCmd.Flags().String("resource-arns", "", "The ARN that identifies the Amazon Aurora or RDS database.")
	})
	computeOptimizerCmd.AddCommand(computeOptimizer_getRdsdatabaseRecommendationsCmd)
}
