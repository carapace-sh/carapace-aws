package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var savingsplans_createSavingsPlanCmd = &cobra.Command{
	Use:   "create-savings-plan",
	Short: "Creates a Savings Plan.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(savingsplans_createSavingsPlanCmd).Standalone()

	savingsplans_createSavingsPlanCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
	savingsplans_createSavingsPlanCmd.Flags().String("commitment", "", "The hourly commitment, in the same currency of the `savingsPlanOfferingId`.")
	savingsplans_createSavingsPlanCmd.Flags().String("purchase-time", "", "The purchase time of the Savings Plan in UTC format (YYYY-MM-DDTHH:MM:SSZ).")
	savingsplans_createSavingsPlanCmd.Flags().String("savings-plan-offering-id", "", "The ID of the offering.")
	savingsplans_createSavingsPlanCmd.Flags().String("tags", "", "One or more tags.")
	savingsplans_createSavingsPlanCmd.Flags().String("upfront-payment-amount", "", "The up-front payment amount.")
	savingsplans_createSavingsPlanCmd.MarkFlagRequired("commitment")
	savingsplans_createSavingsPlanCmd.MarkFlagRequired("savings-plan-offering-id")
	savingsplansCmd.AddCommand(savingsplans_createSavingsPlanCmd)
}
