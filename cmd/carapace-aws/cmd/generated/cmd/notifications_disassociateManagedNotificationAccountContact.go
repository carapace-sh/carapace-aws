package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var notifications_disassociateManagedNotificationAccountContactCmd = &cobra.Command{
	Use:   "disassociate-managed-notification-account-contact",
	Short: "Disassociates an Account Contact with a particular `ManagedNotificationConfiguration`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(notifications_disassociateManagedNotificationAccountContactCmd).Standalone()

	notifications_disassociateManagedNotificationAccountContactCmd.Flags().String("contact-identifier", "", "The unique value of an Account Contact Type to associate with the `ManagedNotificationConfiguration`.")
	notifications_disassociateManagedNotificationAccountContactCmd.Flags().String("managed-notification-configuration-arn", "", "The Amazon Resource Name (ARN) of the `ManagedNotificationConfiguration` to associate with the Account Contact.")
	notifications_disassociateManagedNotificationAccountContactCmd.MarkFlagRequired("contact-identifier")
	notifications_disassociateManagedNotificationAccountContactCmd.MarkFlagRequired("managed-notification-configuration-arn")
	notificationsCmd.AddCommand(notifications_disassociateManagedNotificationAccountContactCmd)
}
