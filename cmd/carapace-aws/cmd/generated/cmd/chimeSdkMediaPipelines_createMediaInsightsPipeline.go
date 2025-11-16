package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkMediaPipelines_createMediaInsightsPipelineCmd = &cobra.Command{
	Use:   "create-media-insights-pipeline",
	Short: "Creates a media insights pipeline.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkMediaPipelines_createMediaInsightsPipelineCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(chimeSdkMediaPipelines_createMediaInsightsPipelineCmd).Standalone()

		chimeSdkMediaPipelines_createMediaInsightsPipelineCmd.Flags().String("client-request-token", "", "The unique identifier for the media insights pipeline request.")
		chimeSdkMediaPipelines_createMediaInsightsPipelineCmd.Flags().String("kinesis-video-stream-recording-source-runtime-configuration", "", "The runtime configuration for the Kinesis video recording stream source.")
		chimeSdkMediaPipelines_createMediaInsightsPipelineCmd.Flags().String("kinesis-video-stream-source-runtime-configuration", "", "The runtime configuration for the Kinesis video stream source of the media insights pipeline.")
		chimeSdkMediaPipelines_createMediaInsightsPipelineCmd.Flags().String("media-insights-pipeline-configuration-arn", "", "The ARN of the pipeline's configuration.")
		chimeSdkMediaPipelines_createMediaInsightsPipelineCmd.Flags().String("media-insights-runtime-metadata", "", "The runtime metadata for the media insights pipeline.")
		chimeSdkMediaPipelines_createMediaInsightsPipelineCmd.Flags().String("s3-recording-sink-runtime-configuration", "", "The runtime configuration for the S3 recording sink.")
		chimeSdkMediaPipelines_createMediaInsightsPipelineCmd.Flags().String("tags", "", "The tags assigned to the media insights pipeline.")
		chimeSdkMediaPipelines_createMediaInsightsPipelineCmd.MarkFlagRequired("media-insights-pipeline-configuration-arn")
	})
	chimeSdkMediaPipelinesCmd.AddCommand(chimeSdkMediaPipelines_createMediaInsightsPipelineCmd)
}
