package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var devopsGuru_removeNotificationChannelCmd = &cobra.Command{
	Use:   "remove-notification-channel",
	Short: "Removes a notification channel from DevOps Guru.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(devopsGuru_removeNotificationChannelCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(devopsGuru_removeNotificationChannelCmd).Standalone()

		devopsGuru_removeNotificationChannelCmd.Flags().String("id", "", "The ID of the notification channel to be removed.")
		devopsGuru_removeNotificationChannelCmd.MarkFlagRequired("id")
	})
	devopsGuruCmd.AddCommand(devopsGuru_removeNotificationChannelCmd)
}
