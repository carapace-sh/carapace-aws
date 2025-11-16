package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkMediaPipelines_createMediaCapturePipelineCmd = &cobra.Command{
	Use:   "create-media-capture-pipeline",
	Short: "Creates a media pipeline.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkMediaPipelines_createMediaCapturePipelineCmd).Standalone()

	chimeSdkMediaPipelines_createMediaCapturePipelineCmd.Flags().String("chime-sdk-meeting-configuration", "", "The configuration for a specified media pipeline.")
	chimeSdkMediaPipelines_createMediaCapturePipelineCmd.Flags().String("client-request-token", "", "The unique identifier for the client request.")
	chimeSdkMediaPipelines_createMediaCapturePipelineCmd.Flags().String("sink-arn", "", "The ARN of the sink type.")
	chimeSdkMediaPipelines_createMediaCapturePipelineCmd.Flags().String("sink-iam-role-arn", "", "The Amazon Resource Name (ARN) of the sink role to be used with `AwsKmsKeyId` in `SseAwsKeyManagementParams`.")
	chimeSdkMediaPipelines_createMediaCapturePipelineCmd.Flags().String("sink-type", "", "Destination type to which the media artifacts are saved.")
	chimeSdkMediaPipelines_createMediaCapturePipelineCmd.Flags().String("source-arn", "", "ARN of the source from which the media artifacts are captured.")
	chimeSdkMediaPipelines_createMediaCapturePipelineCmd.Flags().String("source-type", "", "Source type from which the media artifacts are captured.")
	chimeSdkMediaPipelines_createMediaCapturePipelineCmd.Flags().String("sse-aws-key-management-params", "", "An object that contains server side encryption parameters to be used by media capture pipeline.")
	chimeSdkMediaPipelines_createMediaCapturePipelineCmd.Flags().String("tags", "", "The tag key-value pairs.")
	chimeSdkMediaPipelines_createMediaCapturePipelineCmd.MarkFlagRequired("sink-arn")
	chimeSdkMediaPipelines_createMediaCapturePipelineCmd.MarkFlagRequired("sink-type")
	chimeSdkMediaPipelines_createMediaCapturePipelineCmd.MarkFlagRequired("source-arn")
	chimeSdkMediaPipelines_createMediaCapturePipelineCmd.MarkFlagRequired("source-type")
	chimeSdkMediaPipelinesCmd.AddCommand(chimeSdkMediaPipelines_createMediaCapturePipelineCmd)
}
