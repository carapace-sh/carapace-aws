package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mturk_updateNotificationSettingsCmd = &cobra.Command{
	Use:   "update-notification-settings",
	Short: "The `UpdateNotificationSettings` operation creates, updates, disables or re-enables notifications for a HIT type.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mturk_updateNotificationSettingsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mturk_updateNotificationSettingsCmd).Standalone()

		mturk_updateNotificationSettingsCmd.Flags().Bool("active", false, "Specifies whether notifications are sent for HITs of this HIT type, according to the notification specification.")
		mturk_updateNotificationSettingsCmd.Flags().String("hittype-id", "", "The ID of the HIT type whose notification specification is being updated.")
		mturk_updateNotificationSettingsCmd.Flags().Bool("no-active", false, "Specifies whether notifications are sent for HITs of this HIT type, according to the notification specification.")
		mturk_updateNotificationSettingsCmd.Flags().String("notification", "", "The notification specification for the HIT type.")
		mturk_updateNotificationSettingsCmd.MarkFlagRequired("hittype-id")
		mturk_updateNotificationSettingsCmd.Flag("no-active").Hidden = true
	})
	mturkCmd.AddCommand(mturk_updateNotificationSettingsCmd)
}
