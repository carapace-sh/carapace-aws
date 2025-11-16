package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpoint_updateVoiceChannelCmd = &cobra.Command{
	Use:   "update-voice-channel",
	Short: "Enables the voice channel for an application or updates the status and settings of the voice channel for an application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpoint_updateVoiceChannelCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(pinpoint_updateVoiceChannelCmd).Standalone()

		pinpoint_updateVoiceChannelCmd.Flags().String("application-id", "", "The unique identifier for the application.")
		pinpoint_updateVoiceChannelCmd.Flags().String("voice-channel-request", "", "")
		pinpoint_updateVoiceChannelCmd.MarkFlagRequired("application-id")
		pinpoint_updateVoiceChannelCmd.MarkFlagRequired("voice-channel-request")
	})
	pinpointCmd.AddCommand(pinpoint_updateVoiceChannelCmd)
}
