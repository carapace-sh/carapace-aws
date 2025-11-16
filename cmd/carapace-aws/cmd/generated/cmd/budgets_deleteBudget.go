package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var budgets_deleteBudgetCmd = &cobra.Command{
	Use:   "delete-budget",
	Short: "Deletes a budget.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(budgets_deleteBudgetCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(budgets_deleteBudgetCmd).Standalone()

		budgets_deleteBudgetCmd.Flags().String("account-id", "", "The `accountId` that is associated with the budget that you want to delete.")
		budgets_deleteBudgetCmd.Flags().String("budget-name", "", "The name of the budget that you want to delete.")
		budgets_deleteBudgetCmd.MarkFlagRequired("account-id")
		budgets_deleteBudgetCmd.MarkFlagRequired("budget-name")
	})
	budgetsCmd.AddCommand(budgets_deleteBudgetCmd)
}
