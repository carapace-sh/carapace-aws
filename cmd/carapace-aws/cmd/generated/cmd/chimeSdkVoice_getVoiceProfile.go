package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkVoice_getVoiceProfileCmd = &cobra.Command{
	Use:   "get-voice-profile",
	Short: "Retrieves the details of the specified voice profile.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkVoice_getVoiceProfileCmd).Standalone()

	chimeSdkVoice_getVoiceProfileCmd.Flags().String("voice-profile-id", "", "The voice profile ID.")
	chimeSdkVoice_getVoiceProfileCmd.MarkFlagRequired("voice-profile-id")
	chimeSdkVoiceCmd.AddCommand(chimeSdkVoice_getVoiceProfileCmd)
}
