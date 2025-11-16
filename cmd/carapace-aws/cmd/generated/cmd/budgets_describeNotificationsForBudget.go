package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var budgets_describeNotificationsForBudgetCmd = &cobra.Command{
	Use:   "describe-notifications-for-budget",
	Short: "Lists the notifications that are associated with a budget.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(budgets_describeNotificationsForBudgetCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(budgets_describeNotificationsForBudgetCmd).Standalone()

		budgets_describeNotificationsForBudgetCmd.Flags().String("account-id", "", "The `accountId` that is associated with the budget whose notifications you want descriptions of.")
		budgets_describeNotificationsForBudgetCmd.Flags().String("budget-name", "", "The name of the budget whose notifications you want descriptions of.")
		budgets_describeNotificationsForBudgetCmd.Flags().String("max-results", "", "An optional integer that represents how many entries a paginated response contains.")
		budgets_describeNotificationsForBudgetCmd.Flags().String("next-token", "", "The pagination token that you include in your request to indicate the next set of results that you want to retrieve.")
		budgets_describeNotificationsForBudgetCmd.MarkFlagRequired("account-id")
		budgets_describeNotificationsForBudgetCmd.MarkFlagRequired("budget-name")
	})
	budgetsCmd.AddCommand(budgets_describeNotificationsForBudgetCmd)
}
