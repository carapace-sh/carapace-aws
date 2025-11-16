package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkMediaPipelines_deleteMediaCapturePipelineCmd = &cobra.Command{
	Use:   "delete-media-capture-pipeline",
	Short: "Deletes the media pipeline.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkMediaPipelines_deleteMediaCapturePipelineCmd).Standalone()

	chimeSdkMediaPipelines_deleteMediaCapturePipelineCmd.Flags().String("media-pipeline-id", "", "The ID of the media pipeline being deleted.")
	chimeSdkMediaPipelines_deleteMediaCapturePipelineCmd.MarkFlagRequired("media-pipeline-id")
	chimeSdkMediaPipelinesCmd.AddCommand(chimeSdkMediaPipelines_deleteMediaCapturePipelineCmd)
}
