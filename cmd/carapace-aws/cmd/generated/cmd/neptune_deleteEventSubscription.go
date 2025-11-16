package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var neptune_deleteEventSubscriptionCmd = &cobra.Command{
	Use:   "delete-event-subscription",
	Short: "Deletes an event notification subscription.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(neptune_deleteEventSubscriptionCmd).Standalone()

	neptune_deleteEventSubscriptionCmd.Flags().String("subscription-name", "", "The name of the event notification subscription you want to delete.")
	neptune_deleteEventSubscriptionCmd.MarkFlagRequired("subscription-name")
	neptuneCmd.AddCommand(neptune_deleteEventSubscriptionCmd)
}
