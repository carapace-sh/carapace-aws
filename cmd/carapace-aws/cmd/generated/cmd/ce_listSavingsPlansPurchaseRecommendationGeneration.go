package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ce_listSavingsPlansPurchaseRecommendationGenerationCmd = &cobra.Command{
	Use:   "list-savings-plans-purchase-recommendation-generation",
	Short: "Retrieves a list of your historical recommendation generations within the past 30 days.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ce_listSavingsPlansPurchaseRecommendationGenerationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ce_listSavingsPlansPurchaseRecommendationGenerationCmd).Standalone()

		ce_listSavingsPlansPurchaseRecommendationGenerationCmd.Flags().String("generation-status", "", "The status of the recommendation generation.")
		ce_listSavingsPlansPurchaseRecommendationGenerationCmd.Flags().String("next-page-token", "", "The token to retrieve the next set of results.")
		ce_listSavingsPlansPurchaseRecommendationGenerationCmd.Flags().String("page-size", "", "The number of recommendations that you want returned in a single response object.")
		ce_listSavingsPlansPurchaseRecommendationGenerationCmd.Flags().String("recommendation-ids", "", "The IDs for each specific recommendation.")
	})
	ceCmd.AddCommand(ce_listSavingsPlansPurchaseRecommendationGenerationCmd)
}
