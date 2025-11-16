package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var budgets_describeBudgetNotificationsForAccountCmd = &cobra.Command{
	Use:   "describe-budget-notifications-for-account",
	Short: "Lists the budget names and notifications that are associated with an account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(budgets_describeBudgetNotificationsForAccountCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(budgets_describeBudgetNotificationsForAccountCmd).Standalone()

		budgets_describeBudgetNotificationsForAccountCmd.Flags().String("account-id", "", "")
		budgets_describeBudgetNotificationsForAccountCmd.Flags().String("max-results", "", "An integer that represents how many budgets a paginated response contains.")
		budgets_describeBudgetNotificationsForAccountCmd.Flags().String("next-token", "", "")
		budgets_describeBudgetNotificationsForAccountCmd.MarkFlagRequired("account-id")
	})
	budgetsCmd.AddCommand(budgets_describeBudgetNotificationsForAccountCmd)
}
