package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var devopsGuru_addNotificationChannelCmd = &cobra.Command{
	Use:   "add-notification-channel",
	Short: "Adds a notification channel to DevOps Guru.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(devopsGuru_addNotificationChannelCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(devopsGuru_addNotificationChannelCmd).Standalone()

		devopsGuru_addNotificationChannelCmd.Flags().String("config", "", "A `NotificationChannelConfig` object that specifies what type of notification channel to add.")
		devopsGuru_addNotificationChannelCmd.MarkFlagRequired("config")
	})
	devopsGuruCmd.AddCommand(devopsGuru_addNotificationChannelCmd)
}
