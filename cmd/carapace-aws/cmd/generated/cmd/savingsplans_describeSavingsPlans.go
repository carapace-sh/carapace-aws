package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var savingsplans_describeSavingsPlansCmd = &cobra.Command{
	Use:   "describe-savings-plans",
	Short: "Describes the specified Savings Plans.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(savingsplans_describeSavingsPlansCmd).Standalone()

	savingsplans_describeSavingsPlansCmd.Flags().String("filters", "", "The filters.")
	savingsplans_describeSavingsPlansCmd.Flags().String("max-results", "", "The maximum number of results to return with a single call.")
	savingsplans_describeSavingsPlansCmd.Flags().String("next-token", "", "The token for the next page of results.")
	savingsplans_describeSavingsPlansCmd.Flags().String("savings-plan-arns", "", "The Amazon Resource Names (ARN) of the Savings Plans.")
	savingsplans_describeSavingsPlansCmd.Flags().String("savings-plan-ids", "", "The IDs of the Savings Plans.")
	savingsplans_describeSavingsPlansCmd.Flags().String("states", "", "The current states of the Savings Plans.")
	savingsplansCmd.AddCommand(savingsplans_describeSavingsPlansCmd)
}
