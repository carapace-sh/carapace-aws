package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkVoice_stopVoiceToneAnalysisTaskCmd = &cobra.Command{
	Use:   "stop-voice-tone-analysis-task",
	Short: "Stops a voice tone analysis task.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkVoice_stopVoiceToneAnalysisTaskCmd).Standalone()

	chimeSdkVoice_stopVoiceToneAnalysisTaskCmd.Flags().String("voice-connector-id", "", "The Voice Connector ID.")
	chimeSdkVoice_stopVoiceToneAnalysisTaskCmd.Flags().String("voice-tone-analysis-task-id", "", "The ID of the voice tone analysis task.")
	chimeSdkVoice_stopVoiceToneAnalysisTaskCmd.MarkFlagRequired("voice-connector-id")
	chimeSdkVoice_stopVoiceToneAnalysisTaskCmd.MarkFlagRequired("voice-tone-analysis-task-id")
	chimeSdkVoiceCmd.AddCommand(chimeSdkVoice_stopVoiceToneAnalysisTaskCmd)
}
