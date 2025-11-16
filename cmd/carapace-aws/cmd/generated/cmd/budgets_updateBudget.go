package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var budgets_updateBudgetCmd = &cobra.Command{
	Use:   "update-budget",
	Short: "Updates a budget.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(budgets_updateBudgetCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(budgets_updateBudgetCmd).Standalone()

		budgets_updateBudgetCmd.Flags().String("account-id", "", "The `accountId` that is associated with the budget that you want to update.")
		budgets_updateBudgetCmd.Flags().String("new-budget", "", "The budget that you want to update your budget to.")
		budgets_updateBudgetCmd.MarkFlagRequired("account-id")
		budgets_updateBudgetCmd.MarkFlagRequired("new-budget")
	})
	budgetsCmd.AddCommand(budgets_updateBudgetCmd)
}
