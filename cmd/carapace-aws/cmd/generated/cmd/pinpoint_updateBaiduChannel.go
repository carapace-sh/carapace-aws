package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpoint_updateBaiduChannelCmd = &cobra.Command{
	Use:   "update-baidu-channel",
	Short: "Enables the Baidu channel for an application or updates the status and settings of the Baidu channel for an application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpoint_updateBaiduChannelCmd).Standalone()

	pinpoint_updateBaiduChannelCmd.Flags().String("application-id", "", "The unique identifier for the application.")
	pinpoint_updateBaiduChannelCmd.Flags().String("baidu-channel-request", "", "")
	pinpoint_updateBaiduChannelCmd.MarkFlagRequired("application-id")
	pinpoint_updateBaiduChannelCmd.MarkFlagRequired("baidu-channel-request")
	pinpointCmd.AddCommand(pinpoint_updateBaiduChannelCmd)
}
