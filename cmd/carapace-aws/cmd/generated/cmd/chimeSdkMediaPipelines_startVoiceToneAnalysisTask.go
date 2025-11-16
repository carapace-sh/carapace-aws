package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkMediaPipelines_startVoiceToneAnalysisTaskCmd = &cobra.Command{
	Use:   "start-voice-tone-analysis-task",
	Short: "Starts a voice tone analysis task.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkMediaPipelines_startVoiceToneAnalysisTaskCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(chimeSdkMediaPipelines_startVoiceToneAnalysisTaskCmd).Standalone()

		chimeSdkMediaPipelines_startVoiceToneAnalysisTaskCmd.Flags().String("client-request-token", "", "The unique identifier for the client request.")
		chimeSdkMediaPipelines_startVoiceToneAnalysisTaskCmd.Flags().String("identifier", "", "The unique identifier of the resource to be updated.")
		chimeSdkMediaPipelines_startVoiceToneAnalysisTaskCmd.Flags().String("kinesis-video-stream-source-task-configuration", "", "The task configuration for the Kinesis video stream source of the media insights pipeline.")
		chimeSdkMediaPipelines_startVoiceToneAnalysisTaskCmd.Flags().String("language-code", "", "The language code.")
		chimeSdkMediaPipelines_startVoiceToneAnalysisTaskCmd.MarkFlagRequired("identifier")
		chimeSdkMediaPipelines_startVoiceToneAnalysisTaskCmd.MarkFlagRequired("language-code")
	})
	chimeSdkMediaPipelinesCmd.AddCommand(chimeSdkMediaPipelines_startVoiceToneAnalysisTaskCmd)
}
