package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkMediaPipelines_getMediaPipelineCmd = &cobra.Command{
	Use:   "get-media-pipeline",
	Short: "Gets an existing media pipeline.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkMediaPipelines_getMediaPipelineCmd).Standalone()

	chimeSdkMediaPipelines_getMediaPipelineCmd.Flags().String("media-pipeline-id", "", "The ID of the pipeline that you want to get.")
	chimeSdkMediaPipelines_getMediaPipelineCmd.MarkFlagRequired("media-pipeline-id")
	chimeSdkMediaPipelinesCmd.AddCommand(chimeSdkMediaPipelines_getMediaPipelineCmd)
}
