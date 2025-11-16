package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkMediaPipelines_stopVoiceToneAnalysisTaskCmd = &cobra.Command{
	Use:   "stop-voice-tone-analysis-task",
	Short: "Stops a voice tone analysis task.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkMediaPipelines_stopVoiceToneAnalysisTaskCmd).Standalone()

	chimeSdkMediaPipelines_stopVoiceToneAnalysisTaskCmd.Flags().String("identifier", "", "The unique identifier of the resource to be updated.")
	chimeSdkMediaPipelines_stopVoiceToneAnalysisTaskCmd.Flags().String("voice-tone-analysis-task-id", "", "The ID of the voice tone analysis task.")
	chimeSdkMediaPipelines_stopVoiceToneAnalysisTaskCmd.MarkFlagRequired("identifier")
	chimeSdkMediaPipelines_stopVoiceToneAnalysisTaskCmd.MarkFlagRequired("voice-tone-analysis-task-id")
	chimeSdkMediaPipelinesCmd.AddCommand(chimeSdkMediaPipelines_stopVoiceToneAnalysisTaskCmd)
}
