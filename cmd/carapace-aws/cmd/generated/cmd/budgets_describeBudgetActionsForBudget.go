package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var budgets_describeBudgetActionsForBudgetCmd = &cobra.Command{
	Use:   "describe-budget-actions-for-budget",
	Short: "Describes all of the budget actions for a budget.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(budgets_describeBudgetActionsForBudgetCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(budgets_describeBudgetActionsForBudgetCmd).Standalone()

		budgets_describeBudgetActionsForBudgetCmd.Flags().String("account-id", "", "")
		budgets_describeBudgetActionsForBudgetCmd.Flags().String("budget-name", "", "")
		budgets_describeBudgetActionsForBudgetCmd.Flags().String("max-results", "", "")
		budgets_describeBudgetActionsForBudgetCmd.Flags().String("next-token", "", "")
		budgets_describeBudgetActionsForBudgetCmd.MarkFlagRequired("account-id")
		budgets_describeBudgetActionsForBudgetCmd.MarkFlagRequired("budget-name")
	})
	budgetsCmd.AddCommand(budgets_describeBudgetActionsForBudgetCmd)
}
