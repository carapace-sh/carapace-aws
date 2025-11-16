package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpoint_deleteVoiceChannelCmd = &cobra.Command{
	Use:   "delete-voice-channel",
	Short: "Disables the voice channel for an application and deletes any existing settings for the channel.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpoint_deleteVoiceChannelCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(pinpoint_deleteVoiceChannelCmd).Standalone()

		pinpoint_deleteVoiceChannelCmd.Flags().String("application-id", "", "The unique identifier for the application.")
		pinpoint_deleteVoiceChannelCmd.MarkFlagRequired("application-id")
	})
	pinpointCmd.AddCommand(pinpoint_deleteVoiceChannelCmd)
}
