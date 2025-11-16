package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var budgets_createNotificationCmd = &cobra.Command{
	Use:   "create-notification",
	Short: "Creates a notification.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(budgets_createNotificationCmd).Standalone()

	budgets_createNotificationCmd.Flags().String("account-id", "", "The `accountId` that is associated with the budget that you want to create a notification for.")
	budgets_createNotificationCmd.Flags().String("budget-name", "", "The name of the budget that you want Amazon Web Services to notify you about.")
	budgets_createNotificationCmd.Flags().String("notification", "", "The notification that you want to create.")
	budgets_createNotificationCmd.Flags().String("subscribers", "", "A list of subscribers that you want to associate with the notification.")
	budgets_createNotificationCmd.MarkFlagRequired("account-id")
	budgets_createNotificationCmd.MarkFlagRequired("budget-name")
	budgets_createNotificationCmd.MarkFlagRequired("notification")
	budgets_createNotificationCmd.MarkFlagRequired("subscribers")
	budgetsCmd.AddCommand(budgets_createNotificationCmd)
}
