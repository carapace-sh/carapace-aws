package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var budgets_createBudgetActionCmd = &cobra.Command{
	Use:   "create-budget-action",
	Short: "Creates a budget action.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(budgets_createBudgetActionCmd).Standalone()

	budgets_createBudgetActionCmd.Flags().String("account-id", "", "")
	budgets_createBudgetActionCmd.Flags().String("action-threshold", "", "")
	budgets_createBudgetActionCmd.Flags().String("action-type", "", "The type of action.")
	budgets_createBudgetActionCmd.Flags().String("approval-model", "", "This specifies if the action needs manual or automatic approval.")
	budgets_createBudgetActionCmd.Flags().String("budget-name", "", "")
	budgets_createBudgetActionCmd.Flags().String("definition", "", "")
	budgets_createBudgetActionCmd.Flags().String("execution-role-arn", "", "The role passed for action execution and reversion.")
	budgets_createBudgetActionCmd.Flags().String("notification-type", "", "")
	budgets_createBudgetActionCmd.Flags().String("resource-tags", "", "An optional list of tags to associate with the specified budget action.")
	budgets_createBudgetActionCmd.Flags().String("subscribers", "", "")
	budgets_createBudgetActionCmd.MarkFlagRequired("account-id")
	budgets_createBudgetActionCmd.MarkFlagRequired("action-threshold")
	budgets_createBudgetActionCmd.MarkFlagRequired("action-type")
	budgets_createBudgetActionCmd.MarkFlagRequired("approval-model")
	budgets_createBudgetActionCmd.MarkFlagRequired("budget-name")
	budgets_createBudgetActionCmd.MarkFlagRequired("definition")
	budgets_createBudgetActionCmd.MarkFlagRequired("execution-role-arn")
	budgets_createBudgetActionCmd.MarkFlagRequired("notification-type")
	budgets_createBudgetActionCmd.MarkFlagRequired("subscribers")
	budgetsCmd.AddCommand(budgets_createBudgetActionCmd)
}
