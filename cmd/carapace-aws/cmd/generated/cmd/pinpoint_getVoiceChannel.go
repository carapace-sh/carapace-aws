package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpoint_getVoiceChannelCmd = &cobra.Command{
	Use:   "get-voice-channel",
	Short: "Retrieves information about the status and settings of the voice channel for an application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpoint_getVoiceChannelCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(pinpoint_getVoiceChannelCmd).Standalone()

		pinpoint_getVoiceChannelCmd.Flags().String("application-id", "", "The unique identifier for the application.")
		pinpoint_getVoiceChannelCmd.MarkFlagRequired("application-id")
	})
	pinpointCmd.AddCommand(pinpoint_getVoiceChannelCmd)
}
