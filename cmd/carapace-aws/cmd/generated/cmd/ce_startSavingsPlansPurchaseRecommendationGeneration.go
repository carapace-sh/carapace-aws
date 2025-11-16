package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ce_startSavingsPlansPurchaseRecommendationGenerationCmd = &cobra.Command{
	Use:   "start-savings-plans-purchase-recommendation-generation",
	Short: "Requests a Savings Plans recommendation generation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ce_startSavingsPlansPurchaseRecommendationGenerationCmd).Standalone()

	ceCmd.AddCommand(ce_startSavingsPlansPurchaseRecommendationGenerationCmd)
}
