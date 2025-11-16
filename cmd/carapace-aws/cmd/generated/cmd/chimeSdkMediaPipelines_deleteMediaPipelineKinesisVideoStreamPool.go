package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkMediaPipelines_deleteMediaPipelineKinesisVideoStreamPoolCmd = &cobra.Command{
	Use:   "delete-media-pipeline-kinesis-video-stream-pool",
	Short: "Deletes an Amazon Kinesis Video Stream pool.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkMediaPipelines_deleteMediaPipelineKinesisVideoStreamPoolCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(chimeSdkMediaPipelines_deleteMediaPipelineKinesisVideoStreamPoolCmd).Standalone()

		chimeSdkMediaPipelines_deleteMediaPipelineKinesisVideoStreamPoolCmd.Flags().String("identifier", "", "The unique identifier of the requested resource.")
		chimeSdkMediaPipelines_deleteMediaPipelineKinesisVideoStreamPoolCmd.MarkFlagRequired("identifier")
	})
	chimeSdkMediaPipelinesCmd.AddCommand(chimeSdkMediaPipelines_deleteMediaPipelineKinesisVideoStreamPoolCmd)
}
