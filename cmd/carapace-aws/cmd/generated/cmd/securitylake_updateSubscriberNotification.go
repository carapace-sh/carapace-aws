package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var securitylake_updateSubscriberNotificationCmd = &cobra.Command{
	Use:   "update-subscriber-notification",
	Short: "Updates an existing notification method for the subscription (SQS or HTTPs endpoint) or switches the notification subscription endpoint for a subscriber.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(securitylake_updateSubscriberNotificationCmd).Standalone()

	securitylake_updateSubscriberNotificationCmd.Flags().String("configuration", "", "The configuration for subscriber notification.")
	securitylake_updateSubscriberNotificationCmd.Flags().String("subscriber-id", "", "The subscription ID for which the subscription notification is specified.")
	securitylake_updateSubscriberNotificationCmd.MarkFlagRequired("configuration")
	securitylake_updateSubscriberNotificationCmd.MarkFlagRequired("subscriber-id")
	securitylakeCmd.AddCommand(securitylake_updateSubscriberNotificationCmd)
}
