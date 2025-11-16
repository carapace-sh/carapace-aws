package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var budgets_updateBudgetActionCmd = &cobra.Command{
	Use:   "update-budget-action",
	Short: "Updates a budget action.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(budgets_updateBudgetActionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(budgets_updateBudgetActionCmd).Standalone()

		budgets_updateBudgetActionCmd.Flags().String("account-id", "", "")
		budgets_updateBudgetActionCmd.Flags().String("action-id", "", "A system-generated universally unique identifier (UUID) for the action.")
		budgets_updateBudgetActionCmd.Flags().String("action-threshold", "", "")
		budgets_updateBudgetActionCmd.Flags().String("approval-model", "", "This specifies if the action needs manual or automatic approval.")
		budgets_updateBudgetActionCmd.Flags().String("budget-name", "", "")
		budgets_updateBudgetActionCmd.Flags().String("definition", "", "")
		budgets_updateBudgetActionCmd.Flags().String("execution-role-arn", "", "The role passed for action execution and reversion.")
		budgets_updateBudgetActionCmd.Flags().String("notification-type", "", "")
		budgets_updateBudgetActionCmd.Flags().String("subscribers", "", "")
		budgets_updateBudgetActionCmd.MarkFlagRequired("account-id")
		budgets_updateBudgetActionCmd.MarkFlagRequired("action-id")
		budgets_updateBudgetActionCmd.MarkFlagRequired("budget-name")
	})
	budgetsCmd.AddCommand(budgets_updateBudgetActionCmd)
}
