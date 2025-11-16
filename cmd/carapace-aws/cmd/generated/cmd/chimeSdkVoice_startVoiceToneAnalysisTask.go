package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkVoice_startVoiceToneAnalysisTaskCmd = &cobra.Command{
	Use:   "start-voice-tone-analysis-task",
	Short: "Starts a voice tone analysis task.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkVoice_startVoiceToneAnalysisTaskCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(chimeSdkVoice_startVoiceToneAnalysisTaskCmd).Standalone()

		chimeSdkVoice_startVoiceToneAnalysisTaskCmd.Flags().String("client-request-token", "", "The unique identifier for the client request.")
		chimeSdkVoice_startVoiceToneAnalysisTaskCmd.Flags().String("language-code", "", "The language code.")
		chimeSdkVoice_startVoiceToneAnalysisTaskCmd.Flags().String("transaction-id", "", "The transaction ID.")
		chimeSdkVoice_startVoiceToneAnalysisTaskCmd.Flags().String("voice-connector-id", "", "The Voice Connector ID.")
		chimeSdkVoice_startVoiceToneAnalysisTaskCmd.MarkFlagRequired("language-code")
		chimeSdkVoice_startVoiceToneAnalysisTaskCmd.MarkFlagRequired("transaction-id")
		chimeSdkVoice_startVoiceToneAnalysisTaskCmd.MarkFlagRequired("voice-connector-id")
	})
	chimeSdkVoiceCmd.AddCommand(chimeSdkVoice_startVoiceToneAnalysisTaskCmd)
}
