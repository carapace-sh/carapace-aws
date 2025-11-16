package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var budgets_describeSubscribersForNotificationCmd = &cobra.Command{
	Use:   "describe-subscribers-for-notification",
	Short: "Lists the subscribers that are associated with a notification.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(budgets_describeSubscribersForNotificationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(budgets_describeSubscribersForNotificationCmd).Standalone()

		budgets_describeSubscribersForNotificationCmd.Flags().String("account-id", "", "The `accountId` that is associated with the budget whose subscribers you want descriptions of.")
		budgets_describeSubscribersForNotificationCmd.Flags().String("budget-name", "", "The name of the budget whose subscribers you want descriptions of.")
		budgets_describeSubscribersForNotificationCmd.Flags().String("max-results", "", "An optional integer that represents how many entries a paginated response contains.")
		budgets_describeSubscribersForNotificationCmd.Flags().String("next-token", "", "The pagination token that you include in your request to indicate the next set of results that you want to retrieve.")
		budgets_describeSubscribersForNotificationCmd.Flags().String("notification", "", "The notification whose subscribers you want to list.")
		budgets_describeSubscribersForNotificationCmd.MarkFlagRequired("account-id")
		budgets_describeSubscribersForNotificationCmd.MarkFlagRequired("budget-name")
		budgets_describeSubscribersForNotificationCmd.MarkFlagRequired("notification")
	})
	budgetsCmd.AddCommand(budgets_describeSubscribersForNotificationCmd)
}
