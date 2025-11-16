package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var budgets_createBudgetCmd = &cobra.Command{
	Use:   "create-budget",
	Short: "Creates a budget and, if included, notifications and subscribers.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(budgets_createBudgetCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(budgets_createBudgetCmd).Standalone()

		budgets_createBudgetCmd.Flags().String("account-id", "", "The `accountId` that is associated with the budget.")
		budgets_createBudgetCmd.Flags().String("budget", "", "The budget object that you want to create.")
		budgets_createBudgetCmd.Flags().String("notifications-with-subscribers", "", "A notification that you want to associate with a budget.")
		budgets_createBudgetCmd.Flags().String("resource-tags", "", "An optional list of tags to associate with the specified budget.")
		budgets_createBudgetCmd.MarkFlagRequired("account-id")
		budgets_createBudgetCmd.MarkFlagRequired("budget")
	})
	budgetsCmd.AddCommand(budgets_createBudgetCmd)
}
