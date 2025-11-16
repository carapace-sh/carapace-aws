package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkVoice_stopSpeakerSearchTaskCmd = &cobra.Command{
	Use:   "stop-speaker-search-task",
	Short: "Stops a speaker search task.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkVoice_stopSpeakerSearchTaskCmd).Standalone()

	chimeSdkVoice_stopSpeakerSearchTaskCmd.Flags().String("speaker-search-task-id", "", "The speaker search task ID.")
	chimeSdkVoice_stopSpeakerSearchTaskCmd.Flags().String("voice-connector-id", "", "The Voice Connector ID.")
	chimeSdkVoice_stopSpeakerSearchTaskCmd.MarkFlagRequired("speaker-search-task-id")
	chimeSdkVoice_stopSpeakerSearchTaskCmd.MarkFlagRequired("voice-connector-id")
	chimeSdkVoiceCmd.AddCommand(chimeSdkVoice_stopSpeakerSearchTaskCmd)
}
