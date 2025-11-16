package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var budgets_deleteSubscriberCmd = &cobra.Command{
	Use:   "delete-subscriber",
	Short: "Deletes a subscriber.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(budgets_deleteSubscriberCmd).Standalone()

	budgets_deleteSubscriberCmd.Flags().String("account-id", "", "The `accountId` that is associated with the budget whose subscriber you want to delete.")
	budgets_deleteSubscriberCmd.Flags().String("budget-name", "", "The name of the budget whose subscriber you want to delete.")
	budgets_deleteSubscriberCmd.Flags().String("notification", "", "The notification whose subscriber you want to delete.")
	budgets_deleteSubscriberCmd.Flags().String("subscriber", "", "The subscriber that you want to delete.")
	budgets_deleteSubscriberCmd.MarkFlagRequired("account-id")
	budgets_deleteSubscriberCmd.MarkFlagRequired("budget-name")
	budgets_deleteSubscriberCmd.MarkFlagRequired("notification")
	budgets_deleteSubscriberCmd.MarkFlagRequired("subscriber")
	budgetsCmd.AddCommand(budgets_deleteSubscriberCmd)
}
