package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ses_setIdentityHeadersInNotificationsEnabledCmd = &cobra.Command{
	Use:   "set-identity-headers-in-notifications-enabled",
	Short: "Given an identity (an email address or a domain), sets whether Amazon SES includes the original email headers in the Amazon Simple Notification Service (Amazon SNS) notifications of a specified type.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ses_setIdentityHeadersInNotificationsEnabledCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ses_setIdentityHeadersInNotificationsEnabledCmd).Standalone()

		ses_setIdentityHeadersInNotificationsEnabledCmd.Flags().String("enabled", "", "Sets whether Amazon SES includes the original email headers in Amazon SNS notifications of the specified notification type.")
		ses_setIdentityHeadersInNotificationsEnabledCmd.Flags().String("identity", "", "The identity for which to enable or disable headers in notifications.")
		ses_setIdentityHeadersInNotificationsEnabledCmd.Flags().String("notification-type", "", "The notification type for which to enable or disable headers in notifications.")
		ses_setIdentityHeadersInNotificationsEnabledCmd.MarkFlagRequired("enabled")
		ses_setIdentityHeadersInNotificationsEnabledCmd.MarkFlagRequired("identity")
		ses_setIdentityHeadersInNotificationsEnabledCmd.MarkFlagRequired("notification-type")
	})
	sesCmd.AddCommand(ses_setIdentityHeadersInNotificationsEnabledCmd)
}
