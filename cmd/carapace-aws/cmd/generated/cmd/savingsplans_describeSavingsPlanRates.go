package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var savingsplans_describeSavingsPlanRatesCmd = &cobra.Command{
	Use:   "describe-savings-plan-rates",
	Short: "Describes the rates for the specified Savings Plan.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(savingsplans_describeSavingsPlanRatesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(savingsplans_describeSavingsPlanRatesCmd).Standalone()

		savingsplans_describeSavingsPlanRatesCmd.Flags().String("filters", "", "The filters.")
		savingsplans_describeSavingsPlanRatesCmd.Flags().String("max-results", "", "The maximum number of results to return with a single call.")
		savingsplans_describeSavingsPlanRatesCmd.Flags().String("next-token", "", "The token for the next page of results.")
		savingsplans_describeSavingsPlanRatesCmd.Flags().String("savings-plan-id", "", "The ID of the Savings Plan.")
		savingsplans_describeSavingsPlanRatesCmd.MarkFlagRequired("savings-plan-id")
	})
	savingsplansCmd.AddCommand(savingsplans_describeSavingsPlanRatesCmd)
}
