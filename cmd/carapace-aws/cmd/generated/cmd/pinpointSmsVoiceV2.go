package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpointSmsVoiceV2Cmd = &cobra.Command{
	Use:   "pinpoint-sms-voice-v2",
	Short: "Welcome to the *End User MessagingSMS, version 2 API Reference*.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpointSmsVoiceV2Cmd).Standalone()

	rootCmd.AddCommand(pinpointSmsVoiceV2Cmd)
}
