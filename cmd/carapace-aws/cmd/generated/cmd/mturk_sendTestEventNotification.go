package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mturk_sendTestEventNotificationCmd = &cobra.Command{
	Use:   "send-test-event-notification",
	Short: "The `SendTestEventNotification` operation causes Amazon Mechanical Turk to send a notification message as if a HIT event occurred, according to the provided notification specification.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mturk_sendTestEventNotificationCmd).Standalone()

	mturk_sendTestEventNotificationCmd.Flags().String("notification", "", "The notification specification to test.")
	mturk_sendTestEventNotificationCmd.Flags().String("test-event-type", "", "The event to simulate to test the notification specification.")
	mturk_sendTestEventNotificationCmd.MarkFlagRequired("notification")
	mturk_sendTestEventNotificationCmd.MarkFlagRequired("test-event-type")
	mturkCmd.AddCommand(mturk_sendTestEventNotificationCmd)
}
