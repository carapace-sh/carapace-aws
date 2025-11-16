package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkMediaPipelines_stopSpeakerSearchTaskCmd = &cobra.Command{
	Use:   "stop-speaker-search-task",
	Short: "Stops a speaker search task.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkMediaPipelines_stopSpeakerSearchTaskCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(chimeSdkMediaPipelines_stopSpeakerSearchTaskCmd).Standalone()

		chimeSdkMediaPipelines_stopSpeakerSearchTaskCmd.Flags().String("identifier", "", "The unique identifier of the resource to be updated.")
		chimeSdkMediaPipelines_stopSpeakerSearchTaskCmd.Flags().String("speaker-search-task-id", "", "The speaker search task ID.")
		chimeSdkMediaPipelines_stopSpeakerSearchTaskCmd.MarkFlagRequired("identifier")
		chimeSdkMediaPipelines_stopSpeakerSearchTaskCmd.MarkFlagRequired("speaker-search-task-id")
	})
	chimeSdkMediaPipelinesCmd.AddCommand(chimeSdkMediaPipelines_stopSpeakerSearchTaskCmd)
}
