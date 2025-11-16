package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var budgets_describeBudgetActionsForAccountCmd = &cobra.Command{
	Use:   "describe-budget-actions-for-account",
	Short: "Describes all of the budget actions for an account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(budgets_describeBudgetActionsForAccountCmd).Standalone()

	budgets_describeBudgetActionsForAccountCmd.Flags().String("account-id", "", "")
	budgets_describeBudgetActionsForAccountCmd.Flags().String("max-results", "", "")
	budgets_describeBudgetActionsForAccountCmd.Flags().String("next-token", "", "")
	budgets_describeBudgetActionsForAccountCmd.MarkFlagRequired("account-id")
	budgetsCmd.AddCommand(budgets_describeBudgetActionsForAccountCmd)
}
