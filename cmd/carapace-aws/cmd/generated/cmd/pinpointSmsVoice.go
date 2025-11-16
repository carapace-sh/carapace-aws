package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpointSmsVoiceCmd = &cobra.Command{
	Use:   "pinpoint-sms-voice",
	Short: "Pinpoint SMS and Voice Messaging public facing APIs",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpointSmsVoiceCmd).Standalone()

	rootCmd.AddCommand(pinpointSmsVoiceCmd)
}
