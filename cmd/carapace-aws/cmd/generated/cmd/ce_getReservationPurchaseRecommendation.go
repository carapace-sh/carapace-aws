package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ce_getReservationPurchaseRecommendationCmd = &cobra.Command{
	Use:   "get-reservation-purchase-recommendation",
	Short: "Gets recommendations for reservation purchases.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ce_getReservationPurchaseRecommendationCmd).Standalone()

	ce_getReservationPurchaseRecommendationCmd.Flags().String("account-id", "", "The account ID that's associated with the recommendation.")
	ce_getReservationPurchaseRecommendationCmd.Flags().String("account-scope", "", "The account scope that you want your recommendations for.")
	ce_getReservationPurchaseRecommendationCmd.Flags().String("filter", "", "")
	ce_getReservationPurchaseRecommendationCmd.Flags().String("lookback-period-in-days", "", "The number of previous days that you want Amazon Web Services to consider when it calculates your recommendations.")
	ce_getReservationPurchaseRecommendationCmd.Flags().String("next-page-token", "", "The pagination token that indicates the next set of results that you want to retrieve.")
	ce_getReservationPurchaseRecommendationCmd.Flags().String("page-size", "", "The number of recommendations that you want returned in a single response object.")
	ce_getReservationPurchaseRecommendationCmd.Flags().String("payment-option", "", "The reservation purchase option that you want recommendations for.")
	ce_getReservationPurchaseRecommendationCmd.Flags().String("service", "", "The specific service that you want recommendations for.")
	ce_getReservationPurchaseRecommendationCmd.Flags().String("service-specification", "", "The hardware specifications for the service instances that you want recommendations for, such as standard or convertible Amazon EC2 instances.")
	ce_getReservationPurchaseRecommendationCmd.Flags().String("term-in-years", "", "The reservation term that you want recommendations for.")
	ce_getReservationPurchaseRecommendationCmd.MarkFlagRequired("service")
	ceCmd.AddCommand(ce_getReservationPurchaseRecommendationCmd)
}
