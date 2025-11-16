package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkVoice_getVoiceToneAnalysisTaskCmd = &cobra.Command{
	Use:   "get-voice-tone-analysis-task",
	Short: "Retrieves the details of a voice tone analysis task.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkVoice_getVoiceToneAnalysisTaskCmd).Standalone()

	chimeSdkVoice_getVoiceToneAnalysisTaskCmd.Flags().Bool("is-caller", false, "Specifies whether the voice being analyzed is the caller (originator) or the callee (responder).")
	chimeSdkVoice_getVoiceToneAnalysisTaskCmd.Flags().Bool("no-is-caller", false, "Specifies whether the voice being analyzed is the caller (originator) or the callee (responder).")
	chimeSdkVoice_getVoiceToneAnalysisTaskCmd.Flags().String("voice-connector-id", "", "The Voice Connector ID.")
	chimeSdkVoice_getVoiceToneAnalysisTaskCmd.Flags().String("voice-tone-analysis-task-id", "", "The ID of the voice tone analysis task.")
	chimeSdkVoice_getVoiceToneAnalysisTaskCmd.MarkFlagRequired("is-caller")
	chimeSdkVoice_getVoiceToneAnalysisTaskCmd.Flag("no-is-caller").Hidden = true
	chimeSdkVoice_getVoiceToneAnalysisTaskCmd.MarkFlagRequired("no-is-caller")
	chimeSdkVoice_getVoiceToneAnalysisTaskCmd.MarkFlagRequired("voice-connector-id")
	chimeSdkVoice_getVoiceToneAnalysisTaskCmd.MarkFlagRequired("voice-tone-analysis-task-id")
	chimeSdkVoiceCmd.AddCommand(chimeSdkVoice_getVoiceToneAnalysisTaskCmd)
}
