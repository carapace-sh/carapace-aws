package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkVoiceCmd = &cobra.Command{
	Use:   "chime-sdk-voice",
	Short: "The Amazon Chime SDK telephony APIs in this section enable developers to create PSTN calling solutions that use Amazon Chime SDK Voice Connectors, and Amazon Chime SDK SIP media applications.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkVoiceCmd).Standalone()

	rootCmd.AddCommand(chimeSdkVoiceCmd)
}
