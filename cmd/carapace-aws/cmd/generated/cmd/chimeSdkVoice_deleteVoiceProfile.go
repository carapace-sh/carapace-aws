package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkVoice_deleteVoiceProfileCmd = &cobra.Command{
	Use:   "delete-voice-profile",
	Short: "Deletes a voice profile, including its voice print and enrollment data.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkVoice_deleteVoiceProfileCmd).Standalone()

	chimeSdkVoice_deleteVoiceProfileCmd.Flags().String("voice-profile-id", "", "The voice profile ID.")
	chimeSdkVoice_deleteVoiceProfileCmd.MarkFlagRequired("voice-profile-id")
	chimeSdkVoiceCmd.AddCommand(chimeSdkVoice_deleteVoiceProfileCmd)
}
