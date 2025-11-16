package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ce_getSavingsPlanPurchaseRecommendationDetailsCmd = &cobra.Command{
	Use:   "get-savings-plan-purchase-recommendation-details",
	Short: "Retrieves the details for a Savings Plan recommendation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ce_getSavingsPlanPurchaseRecommendationDetailsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ce_getSavingsPlanPurchaseRecommendationDetailsCmd).Standalone()

		ce_getSavingsPlanPurchaseRecommendationDetailsCmd.Flags().String("recommendation-detail-id", "", "The ID that is associated with the Savings Plan recommendation.")
		ce_getSavingsPlanPurchaseRecommendationDetailsCmd.MarkFlagRequired("recommendation-detail-id")
	})
	ceCmd.AddCommand(ce_getSavingsPlanPurchaseRecommendationDetailsCmd)
}
