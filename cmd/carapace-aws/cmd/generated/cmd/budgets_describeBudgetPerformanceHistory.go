package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var budgets_describeBudgetPerformanceHistoryCmd = &cobra.Command{
	Use:   "describe-budget-performance-history",
	Short: "Describes the history for `DAILY`, `MONTHLY`, and `QUARTERLY` budgets.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(budgets_describeBudgetPerformanceHistoryCmd).Standalone()

	budgets_describeBudgetPerformanceHistoryCmd.Flags().String("account-id", "", "")
	budgets_describeBudgetPerformanceHistoryCmd.Flags().String("budget-name", "", "")
	budgets_describeBudgetPerformanceHistoryCmd.Flags().String("max-results", "", "")
	budgets_describeBudgetPerformanceHistoryCmd.Flags().String("next-token", "", "")
	budgets_describeBudgetPerformanceHistoryCmd.Flags().String("time-period", "", "Retrieves how often the budget went into an `ALARM` state for the specified time period.")
	budgets_describeBudgetPerformanceHistoryCmd.MarkFlagRequired("account-id")
	budgets_describeBudgetPerformanceHistoryCmd.MarkFlagRequired("budget-name")
	budgetsCmd.AddCommand(budgets_describeBudgetPerformanceHistoryCmd)
}
