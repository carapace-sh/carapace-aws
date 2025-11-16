package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rolesanywhere_resetNotificationSettingsCmd = &cobra.Command{
	Use:   "reset-notification-settings",
	Short: "Resets the *custom notification setting* to IAM Roles Anywhere default setting.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rolesanywhere_resetNotificationSettingsCmd).Standalone()

	rolesanywhere_resetNotificationSettingsCmd.Flags().String("notification-setting-keys", "", "A list of notification setting keys to reset.")
	rolesanywhere_resetNotificationSettingsCmd.Flags().String("trust-anchor-id", "", "The unique identifier of the trust anchor.")
	rolesanywhere_resetNotificationSettingsCmd.MarkFlagRequired("notification-setting-keys")
	rolesanywhere_resetNotificationSettingsCmd.MarkFlagRequired("trust-anchor-id")
	rolesanywhereCmd.AddCommand(rolesanywhere_resetNotificationSettingsCmd)
}
