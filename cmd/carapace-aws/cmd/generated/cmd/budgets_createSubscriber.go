package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var budgets_createSubscriberCmd = &cobra.Command{
	Use:   "create-subscriber",
	Short: "Creates a subscriber.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(budgets_createSubscriberCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(budgets_createSubscriberCmd).Standalone()

		budgets_createSubscriberCmd.Flags().String("account-id", "", "The `accountId` that is associated with the budget that you want to create a subscriber for.")
		budgets_createSubscriberCmd.Flags().String("budget-name", "", "The name of the budget that you want to subscribe to.")
		budgets_createSubscriberCmd.Flags().String("notification", "", "The notification that you want to create a subscriber for.")
		budgets_createSubscriberCmd.Flags().String("subscriber", "", "The subscriber that you want to associate with a budget notification.")
		budgets_createSubscriberCmd.MarkFlagRequired("account-id")
		budgets_createSubscriberCmd.MarkFlagRequired("budget-name")
		budgets_createSubscriberCmd.MarkFlagRequired("notification")
		budgets_createSubscriberCmd.MarkFlagRequired("subscriber")
	})
	budgetsCmd.AddCommand(budgets_createSubscriberCmd)
}
