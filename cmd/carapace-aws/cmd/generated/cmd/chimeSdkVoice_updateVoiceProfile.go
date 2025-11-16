package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkVoice_updateVoiceProfileCmd = &cobra.Command{
	Use:   "update-voice-profile",
	Short: "Updates the specified voice profile’s voice print and refreshes its expiration timestamp.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkVoice_updateVoiceProfileCmd).Standalone()

	chimeSdkVoice_updateVoiceProfileCmd.Flags().String("speaker-search-task-id", "", "The ID of the speaker search task.")
	chimeSdkVoice_updateVoiceProfileCmd.Flags().String("voice-profile-id", "", "The profile ID.")
	chimeSdkVoice_updateVoiceProfileCmd.MarkFlagRequired("speaker-search-task-id")
	chimeSdkVoice_updateVoiceProfileCmd.MarkFlagRequired("voice-profile-id")
	chimeSdkVoiceCmd.AddCommand(chimeSdkVoice_updateVoiceProfileCmd)
}
