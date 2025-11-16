package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkMediaPipelines_getMediaCapturePipelineCmd = &cobra.Command{
	Use:   "get-media-capture-pipeline",
	Short: "Gets an existing media pipeline.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkMediaPipelines_getMediaCapturePipelineCmd).Standalone()

	chimeSdkMediaPipelines_getMediaCapturePipelineCmd.Flags().String("media-pipeline-id", "", "The ID of the pipeline that you want to get.")
	chimeSdkMediaPipelines_getMediaCapturePipelineCmd.MarkFlagRequired("media-pipeline-id")
	chimeSdkMediaPipelinesCmd.AddCommand(chimeSdkMediaPipelines_getMediaCapturePipelineCmd)
}
