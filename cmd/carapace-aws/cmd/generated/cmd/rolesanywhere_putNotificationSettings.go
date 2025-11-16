package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rolesanywhere_putNotificationSettingsCmd = &cobra.Command{
	Use:   "put-notification-settings",
	Short: "Attaches a list of *notification settings* to a trust anchor.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rolesanywhere_putNotificationSettingsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(rolesanywhere_putNotificationSettingsCmd).Standalone()

		rolesanywhere_putNotificationSettingsCmd.Flags().String("notification-settings", "", "A list of notification settings to be associated to the trust anchor.")
		rolesanywhere_putNotificationSettingsCmd.Flags().String("trust-anchor-id", "", "The unique identifier of the trust anchor.")
		rolesanywhere_putNotificationSettingsCmd.MarkFlagRequired("notification-settings")
		rolesanywhere_putNotificationSettingsCmd.MarkFlagRequired("trust-anchor-id")
	})
	rolesanywhereCmd.AddCommand(rolesanywhere_putNotificationSettingsCmd)
}
