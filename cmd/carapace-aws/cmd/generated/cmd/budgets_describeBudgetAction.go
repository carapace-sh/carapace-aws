package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var budgets_describeBudgetActionCmd = &cobra.Command{
	Use:   "describe-budget-action",
	Short: "Describes a budget action detail.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(budgets_describeBudgetActionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(budgets_describeBudgetActionCmd).Standalone()

		budgets_describeBudgetActionCmd.Flags().String("account-id", "", "")
		budgets_describeBudgetActionCmd.Flags().String("action-id", "", "A system-generated universally unique identifier (UUID) for the action.")
		budgets_describeBudgetActionCmd.Flags().String("budget-name", "", "")
		budgets_describeBudgetActionCmd.MarkFlagRequired("account-id")
		budgets_describeBudgetActionCmd.MarkFlagRequired("action-id")
		budgets_describeBudgetActionCmd.MarkFlagRequired("budget-name")
	})
	budgetsCmd.AddCommand(budgets_describeBudgetActionCmd)
}
