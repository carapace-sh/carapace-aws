package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkMediaPipelines_createMediaPipelineKinesisVideoStreamPoolCmd = &cobra.Command{
	Use:   "create-media-pipeline-kinesis-video-stream-pool",
	Short: "Creates an Amazon Kinesis Video Stream pool for use with media stream pipelines.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkMediaPipelines_createMediaPipelineKinesisVideoStreamPoolCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(chimeSdkMediaPipelines_createMediaPipelineKinesisVideoStreamPoolCmd).Standalone()

		chimeSdkMediaPipelines_createMediaPipelineKinesisVideoStreamPoolCmd.Flags().String("client-request-token", "", "The token assigned to the client making the request.")
		chimeSdkMediaPipelines_createMediaPipelineKinesisVideoStreamPoolCmd.Flags().String("pool-name", "", "The name of the pool.")
		chimeSdkMediaPipelines_createMediaPipelineKinesisVideoStreamPoolCmd.Flags().String("stream-configuration", "", "The configuration settings for the stream.")
		chimeSdkMediaPipelines_createMediaPipelineKinesisVideoStreamPoolCmd.Flags().String("tags", "", "The tags assigned to the stream pool.")
		chimeSdkMediaPipelines_createMediaPipelineKinesisVideoStreamPoolCmd.MarkFlagRequired("pool-name")
		chimeSdkMediaPipelines_createMediaPipelineKinesisVideoStreamPoolCmd.MarkFlagRequired("stream-configuration")
	})
	chimeSdkMediaPipelinesCmd.AddCommand(chimeSdkMediaPipelines_createMediaPipelineKinesisVideoStreamPoolCmd)
}
