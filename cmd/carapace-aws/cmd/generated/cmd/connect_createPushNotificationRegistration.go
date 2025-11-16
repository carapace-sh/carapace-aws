package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_createPushNotificationRegistrationCmd = &cobra.Command{
	Use:   "create-push-notification-registration",
	Short: "Creates registration for a device token and a chat contact to receive real-time push notifications.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_createPushNotificationRegistrationCmd).Standalone()

	connect_createPushNotificationRegistrationCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
	connect_createPushNotificationRegistrationCmd.Flags().String("contact-configuration", "", "The contact configuration for push notification registration.")
	connect_createPushNotificationRegistrationCmd.Flags().String("device-token", "", "The push notification token issued by the Apple or Google gateways.")
	connect_createPushNotificationRegistrationCmd.Flags().String("device-type", "", "The device type to use when sending the message.")
	connect_createPushNotificationRegistrationCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
	connect_createPushNotificationRegistrationCmd.Flags().String("pinpoint-app-arn", "", "The Amazon Resource Name (ARN) of the Pinpoint application.")
	connect_createPushNotificationRegistrationCmd.MarkFlagRequired("contact-configuration")
	connect_createPushNotificationRegistrationCmd.MarkFlagRequired("device-token")
	connect_createPushNotificationRegistrationCmd.MarkFlagRequired("device-type")
	connect_createPushNotificationRegistrationCmd.MarkFlagRequired("instance-id")
	connect_createPushNotificationRegistrationCmd.MarkFlagRequired("pinpoint-app-arn")
	connectCmd.AddCommand(connect_createPushNotificationRegistrationCmd)
}
