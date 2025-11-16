package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var budgets_updateNotificationCmd = &cobra.Command{
	Use:   "update-notification",
	Short: "Updates a notification.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(budgets_updateNotificationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(budgets_updateNotificationCmd).Standalone()

		budgets_updateNotificationCmd.Flags().String("account-id", "", "The `accountId` that is associated with the budget whose notification you want to update.")
		budgets_updateNotificationCmd.Flags().String("budget-name", "", "The name of the budget whose notification you want to update.")
		budgets_updateNotificationCmd.Flags().String("new-notification", "", "The updated notification to be associated with a budget.")
		budgets_updateNotificationCmd.Flags().String("old-notification", "", "The previous notification that is associated with a budget.")
		budgets_updateNotificationCmd.MarkFlagRequired("account-id")
		budgets_updateNotificationCmd.MarkFlagRequired("budget-name")
		budgets_updateNotificationCmd.MarkFlagRequired("new-notification")
		budgets_updateNotificationCmd.MarkFlagRequired("old-notification")
	})
	budgetsCmd.AddCommand(budgets_updateNotificationCmd)
}
