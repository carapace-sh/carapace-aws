package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var budgets_deleteBudgetActionCmd = &cobra.Command{
	Use:   "delete-budget-action",
	Short: "Deletes a budget action.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(budgets_deleteBudgetActionCmd).Standalone()

	budgets_deleteBudgetActionCmd.Flags().String("account-id", "", "")
	budgets_deleteBudgetActionCmd.Flags().String("action-id", "", "A system-generated universally unique identifier (UUID) for the action.")
	budgets_deleteBudgetActionCmd.Flags().String("budget-name", "", "")
	budgets_deleteBudgetActionCmd.MarkFlagRequired("account-id")
	budgets_deleteBudgetActionCmd.MarkFlagRequired("action-id")
	budgets_deleteBudgetActionCmd.MarkFlagRequired("budget-name")
	budgetsCmd.AddCommand(budgets_deleteBudgetActionCmd)
}
