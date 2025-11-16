package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var budgets_executeBudgetActionCmd = &cobra.Command{
	Use:   "execute-budget-action",
	Short: "Executes a budget action.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(budgets_executeBudgetActionCmd).Standalone()

	budgets_executeBudgetActionCmd.Flags().String("account-id", "", "")
	budgets_executeBudgetActionCmd.Flags().String("action-id", "", "A system-generated universally unique identifier (UUID) for the action.")
	budgets_executeBudgetActionCmd.Flags().String("budget-name", "", "")
	budgets_executeBudgetActionCmd.Flags().String("execution-type", "", "The type of execution.")
	budgets_executeBudgetActionCmd.MarkFlagRequired("account-id")
	budgets_executeBudgetActionCmd.MarkFlagRequired("action-id")
	budgets_executeBudgetActionCmd.MarkFlagRequired("budget-name")
	budgets_executeBudgetActionCmd.MarkFlagRequired("execution-type")
	budgetsCmd.AddCommand(budgets_executeBudgetActionCmd)
}
