package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var budgets_describeBudgetCmd = &cobra.Command{
	Use:   "describe-budget",
	Short: "Describes a budget.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(budgets_describeBudgetCmd).Standalone()

	budgets_describeBudgetCmd.Flags().String("account-id", "", "The `accountId` that is associated with the budget that you want a description of.")
	budgets_describeBudgetCmd.Flags().String("budget-name", "", "The name of the budget that you want a description of.")
	budgets_describeBudgetCmd.Flags().String("show-filter-expression", "", "Specifies whether the response includes the filter expression associated with the budget.")
	budgets_describeBudgetCmd.MarkFlagRequired("account-id")
	budgets_describeBudgetCmd.MarkFlagRequired("budget-name")
	budgetsCmd.AddCommand(budgets_describeBudgetCmd)
}
