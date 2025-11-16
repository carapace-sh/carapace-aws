package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ce_getRightsizingRecommendationCmd = &cobra.Command{
	Use:   "get-rightsizing-recommendation",
	Short: "Creates recommendations that help you save cost by identifying idle and underutilized Amazon EC2 instances.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ce_getRightsizingRecommendationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ce_getRightsizingRecommendationCmd).Standalone()

		ce_getRightsizingRecommendationCmd.Flags().String("configuration", "", "You can use Configuration to customize recommendations across two attributes.")
		ce_getRightsizingRecommendationCmd.Flags().String("filter", "", "")
		ce_getRightsizingRecommendationCmd.Flags().String("next-page-token", "", "The pagination token that indicates the next set of results that you want to retrieve.")
		ce_getRightsizingRecommendationCmd.Flags().String("page-size", "", "The number of recommendations that you want returned in a single response object.")
		ce_getRightsizingRecommendationCmd.Flags().String("service", "", "The specific service that you want recommendations for.")
		ce_getRightsizingRecommendationCmd.MarkFlagRequired("service")
	})
	ceCmd.AddCommand(ce_getRightsizingRecommendationCmd)
}
