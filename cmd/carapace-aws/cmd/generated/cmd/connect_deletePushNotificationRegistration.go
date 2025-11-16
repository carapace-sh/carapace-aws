package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_deletePushNotificationRegistrationCmd = &cobra.Command{
	Use:   "delete-push-notification-registration",
	Short: "Deletes registration for a device token and a chat contact.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_deletePushNotificationRegistrationCmd).Standalone()

	connect_deletePushNotificationRegistrationCmd.Flags().String("contact-id", "", "The identifier of the contact within the Amazon Connect instance.")
	connect_deletePushNotificationRegistrationCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
	connect_deletePushNotificationRegistrationCmd.Flags().String("registration-id", "", "The identifier for the registration.")
	connect_deletePushNotificationRegistrationCmd.MarkFlagRequired("contact-id")
	connect_deletePushNotificationRegistrationCmd.MarkFlagRequired("instance-id")
	connect_deletePushNotificationRegistrationCmd.MarkFlagRequired("registration-id")
	connectCmd.AddCommand(connect_deletePushNotificationRegistrationCmd)
}
