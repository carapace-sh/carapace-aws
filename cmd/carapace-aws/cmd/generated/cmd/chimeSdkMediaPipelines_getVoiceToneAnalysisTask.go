package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkMediaPipelines_getVoiceToneAnalysisTaskCmd = &cobra.Command{
	Use:   "get-voice-tone-analysis-task",
	Short: "Retrieves the details of a voice tone analysis task.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkMediaPipelines_getVoiceToneAnalysisTaskCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(chimeSdkMediaPipelines_getVoiceToneAnalysisTaskCmd).Standalone()

		chimeSdkMediaPipelines_getVoiceToneAnalysisTaskCmd.Flags().String("identifier", "", "The unique identifier of the resource to be updated.")
		chimeSdkMediaPipelines_getVoiceToneAnalysisTaskCmd.Flags().String("voice-tone-analysis-task-id", "", "The ID of the voice tone analysis task.")
		chimeSdkMediaPipelines_getVoiceToneAnalysisTaskCmd.MarkFlagRequired("identifier")
		chimeSdkMediaPipelines_getVoiceToneAnalysisTaskCmd.MarkFlagRequired("voice-tone-analysis-task-id")
	})
	chimeSdkMediaPipelinesCmd.AddCommand(chimeSdkMediaPipelines_getVoiceToneAnalysisTaskCmd)
}
