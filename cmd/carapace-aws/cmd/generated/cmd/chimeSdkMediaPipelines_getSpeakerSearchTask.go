package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkMediaPipelines_getSpeakerSearchTaskCmd = &cobra.Command{
	Use:   "get-speaker-search-task",
	Short: "Retrieves the details of the specified speaker search task.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkMediaPipelines_getSpeakerSearchTaskCmd).Standalone()

	chimeSdkMediaPipelines_getSpeakerSearchTaskCmd.Flags().String("identifier", "", "The unique identifier of the resource to be updated.")
	chimeSdkMediaPipelines_getSpeakerSearchTaskCmd.Flags().String("speaker-search-task-id", "", "The ID of the speaker search task.")
	chimeSdkMediaPipelines_getSpeakerSearchTaskCmd.MarkFlagRequired("identifier")
	chimeSdkMediaPipelines_getSpeakerSearchTaskCmd.MarkFlagRequired("speaker-search-task-id")
	chimeSdkMediaPipelinesCmd.AddCommand(chimeSdkMediaPipelines_getSpeakerSearchTaskCmd)
}
