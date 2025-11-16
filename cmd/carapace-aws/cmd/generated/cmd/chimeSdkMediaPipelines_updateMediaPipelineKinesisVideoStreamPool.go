package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkMediaPipelines_updateMediaPipelineKinesisVideoStreamPoolCmd = &cobra.Command{
	Use:   "update-media-pipeline-kinesis-video-stream-pool",
	Short: "Updates an Amazon Kinesis Video Stream pool in a media pipeline.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkMediaPipelines_updateMediaPipelineKinesisVideoStreamPoolCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(chimeSdkMediaPipelines_updateMediaPipelineKinesisVideoStreamPoolCmd).Standalone()

		chimeSdkMediaPipelines_updateMediaPipelineKinesisVideoStreamPoolCmd.Flags().String("identifier", "", "The unique identifier of the requested resource.")
		chimeSdkMediaPipelines_updateMediaPipelineKinesisVideoStreamPoolCmd.Flags().String("stream-configuration", "", "The configuration settings for the video stream.")
		chimeSdkMediaPipelines_updateMediaPipelineKinesisVideoStreamPoolCmd.MarkFlagRequired("identifier")
	})
	chimeSdkMediaPipelinesCmd.AddCommand(chimeSdkMediaPipelines_updateMediaPipelineKinesisVideoStreamPoolCmd)
}
