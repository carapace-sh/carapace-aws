package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var budgets_deleteNotificationCmd = &cobra.Command{
	Use:   "delete-notification",
	Short: "Deletes a notification.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(budgets_deleteNotificationCmd).Standalone()

	budgets_deleteNotificationCmd.Flags().String("account-id", "", "The `accountId` that is associated with the budget whose notification you want to delete.")
	budgets_deleteNotificationCmd.Flags().String("budget-name", "", "The name of the budget whose notification you want to delete.")
	budgets_deleteNotificationCmd.Flags().String("notification", "", "The notification that you want to delete.")
	budgets_deleteNotificationCmd.MarkFlagRequired("account-id")
	budgets_deleteNotificationCmd.MarkFlagRequired("budget-name")
	budgets_deleteNotificationCmd.MarkFlagRequired("notification")
	budgetsCmd.AddCommand(budgets_deleteNotificationCmd)
}
