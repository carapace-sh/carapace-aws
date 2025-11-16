package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var notifications_associateManagedNotificationAccountContactCmd = &cobra.Command{
	Use:   "associate-managed-notification-account-contact",
	Short: "Associates an Account Contact with a particular `ManagedNotificationConfiguration`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(notifications_associateManagedNotificationAccountContactCmd).Standalone()

	notifications_associateManagedNotificationAccountContactCmd.Flags().String("contact-identifier", "", "A unique value of an Account Contact Type to associate with the `ManagedNotificationConfiguration`.")
	notifications_associateManagedNotificationAccountContactCmd.Flags().String("managed-notification-configuration-arn", "", "The Amazon Resource Name (ARN) of the `ManagedNotificationConfiguration` to associate with the Account Contact.")
	notifications_associateManagedNotificationAccountContactCmd.MarkFlagRequired("contact-identifier")
	notifications_associateManagedNotificationAccountContactCmd.MarkFlagRequired("managed-notification-configuration-arn")
	notificationsCmd.AddCommand(notifications_associateManagedNotificationAccountContactCmd)
}
