package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpoint_deleteBaiduChannelCmd = &cobra.Command{
	Use:   "delete-baidu-channel",
	Short: "Disables the Baidu channel for an application and deletes any existing settings for the channel.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpoint_deleteBaiduChannelCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(pinpoint_deleteBaiduChannelCmd).Standalone()

		pinpoint_deleteBaiduChannelCmd.Flags().String("application-id", "", "The unique identifier for the application.")
		pinpoint_deleteBaiduChannelCmd.MarkFlagRequired("application-id")
	})
	pinpointCmd.AddCommand(pinpoint_deleteBaiduChannelCmd)
}
