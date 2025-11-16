package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkVoice_getSpeakerSearchTaskCmd = &cobra.Command{
	Use:   "get-speaker-search-task",
	Short: "Retrieves the details of the specified speaker search task.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkVoice_getSpeakerSearchTaskCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(chimeSdkVoice_getSpeakerSearchTaskCmd).Standalone()

		chimeSdkVoice_getSpeakerSearchTaskCmd.Flags().String("speaker-search-task-id", "", "The ID of the speaker search task.")
		chimeSdkVoice_getSpeakerSearchTaskCmd.Flags().String("voice-connector-id", "", "The Voice Connector ID.")
		chimeSdkVoice_getSpeakerSearchTaskCmd.MarkFlagRequired("speaker-search-task-id")
		chimeSdkVoice_getSpeakerSearchTaskCmd.MarkFlagRequired("voice-connector-id")
	})
	chimeSdkVoiceCmd.AddCommand(chimeSdkVoice_getSpeakerSearchTaskCmd)
}
