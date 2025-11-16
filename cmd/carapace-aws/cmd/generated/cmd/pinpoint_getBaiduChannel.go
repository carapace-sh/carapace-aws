package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpoint_getBaiduChannelCmd = &cobra.Command{
	Use:   "get-baidu-channel",
	Short: "Retrieves information about the status and settings of the Baidu channel for an application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpoint_getBaiduChannelCmd).Standalone()

	pinpoint_getBaiduChannelCmd.Flags().String("application-id", "", "The unique identifier for the application.")
	pinpoint_getBaiduChannelCmd.MarkFlagRequired("application-id")
	pinpointCmd.AddCommand(pinpoint_getBaiduChannelCmd)
}
