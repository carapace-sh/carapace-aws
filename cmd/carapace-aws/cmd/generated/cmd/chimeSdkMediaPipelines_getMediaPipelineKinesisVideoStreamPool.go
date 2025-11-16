package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkMediaPipelines_getMediaPipelineKinesisVideoStreamPoolCmd = &cobra.Command{
	Use:   "get-media-pipeline-kinesis-video-stream-pool",
	Short: "Gets an Kinesis video stream pool.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkMediaPipelines_getMediaPipelineKinesisVideoStreamPoolCmd).Standalone()

	chimeSdkMediaPipelines_getMediaPipelineKinesisVideoStreamPoolCmd.Flags().String("identifier", "", "The unique identifier of the requested resource.")
	chimeSdkMediaPipelines_getMediaPipelineKinesisVideoStreamPoolCmd.MarkFlagRequired("identifier")
	chimeSdkMediaPipelinesCmd.AddCommand(chimeSdkMediaPipelines_getMediaPipelineKinesisVideoStreamPoolCmd)
}
