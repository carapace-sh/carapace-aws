package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var budgets_updateSubscriberCmd = &cobra.Command{
	Use:   "update-subscriber",
	Short: "Updates a subscriber.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(budgets_updateSubscriberCmd).Standalone()

	budgets_updateSubscriberCmd.Flags().String("account-id", "", "The `accountId` that is associated with the budget whose subscriber you want to update.")
	budgets_updateSubscriberCmd.Flags().String("budget-name", "", "The name of the budget whose subscriber you want to update.")
	budgets_updateSubscriberCmd.Flags().String("new-subscriber", "", "The updated subscriber that is associated with a budget notification.")
	budgets_updateSubscriberCmd.Flags().String("notification", "", "The notification whose subscriber you want to update.")
	budgets_updateSubscriberCmd.Flags().String("old-subscriber", "", "The previous subscriber that is associated with a budget notification.")
	budgets_updateSubscriberCmd.MarkFlagRequired("account-id")
	budgets_updateSubscriberCmd.MarkFlagRequired("budget-name")
	budgets_updateSubscriberCmd.MarkFlagRequired("new-subscriber")
	budgets_updateSubscriberCmd.MarkFlagRequired("notification")
	budgets_updateSubscriberCmd.MarkFlagRequired("old-subscriber")
	budgetsCmd.AddCommand(budgets_updateSubscriberCmd)
}
