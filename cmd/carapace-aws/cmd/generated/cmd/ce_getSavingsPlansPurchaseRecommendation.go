package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ce_getSavingsPlansPurchaseRecommendationCmd = &cobra.Command{
	Use:   "get-savings-plans-purchase-recommendation",
	Short: "Retrieves the Savings Plans recommendations for your account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ce_getSavingsPlansPurchaseRecommendationCmd).Standalone()

	ce_getSavingsPlansPurchaseRecommendationCmd.Flags().String("account-scope", "", "The account scope that you want your recommendations for.")
	ce_getSavingsPlansPurchaseRecommendationCmd.Flags().String("filter", "", "You can filter your recommendations by Account ID with the `LINKED_ACCOUNT` dimension.")
	ce_getSavingsPlansPurchaseRecommendationCmd.Flags().String("lookback-period-in-days", "", "The lookback period that's used to generate the recommendation.")
	ce_getSavingsPlansPurchaseRecommendationCmd.Flags().String("next-page-token", "", "The token to retrieve the next set of results.")
	ce_getSavingsPlansPurchaseRecommendationCmd.Flags().String("page-size", "", "The number of recommendations that you want returned in a single response object.")
	ce_getSavingsPlansPurchaseRecommendationCmd.Flags().String("payment-option", "", "The payment option that's used to generate these recommendations.")
	ce_getSavingsPlansPurchaseRecommendationCmd.Flags().String("savings-plans-type", "", "The Savings Plans recommendation type that's requested.")
	ce_getSavingsPlansPurchaseRecommendationCmd.Flags().String("term-in-years", "", "The savings plan recommendation term that's used to generate these recommendations.")
	ce_getSavingsPlansPurchaseRecommendationCmd.MarkFlagRequired("lookback-period-in-days")
	ce_getSavingsPlansPurchaseRecommendationCmd.MarkFlagRequired("payment-option")
	ce_getSavingsPlansPurchaseRecommendationCmd.MarkFlagRequired("savings-plans-type")
	ce_getSavingsPlansPurchaseRecommendationCmd.MarkFlagRequired("term-in-years")
	ceCmd.AddCommand(ce_getSavingsPlansPurchaseRecommendationCmd)
}
