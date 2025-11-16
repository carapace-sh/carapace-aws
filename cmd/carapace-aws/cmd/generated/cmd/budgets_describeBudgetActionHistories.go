package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var budgets_describeBudgetActionHistoriesCmd = &cobra.Command{
	Use:   "describe-budget-action-histories",
	Short: "Describes a budget action history detail.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(budgets_describeBudgetActionHistoriesCmd).Standalone()

	budgets_describeBudgetActionHistoriesCmd.Flags().String("account-id", "", "")
	budgets_describeBudgetActionHistoriesCmd.Flags().String("action-id", "", "A system-generated universally unique identifier (UUID) for the action.")
	budgets_describeBudgetActionHistoriesCmd.Flags().String("budget-name", "", "")
	budgets_describeBudgetActionHistoriesCmd.Flags().String("max-results", "", "")
	budgets_describeBudgetActionHistoriesCmd.Flags().String("next-token", "", "")
	budgets_describeBudgetActionHistoriesCmd.Flags().String("time-period", "", "")
	budgets_describeBudgetActionHistoriesCmd.MarkFlagRequired("account-id")
	budgets_describeBudgetActionHistoriesCmd.MarkFlagRequired("action-id")
	budgets_describeBudgetActionHistoriesCmd.MarkFlagRequired("budget-name")
	budgetsCmd.AddCommand(budgets_describeBudgetActionHistoriesCmd)
}
