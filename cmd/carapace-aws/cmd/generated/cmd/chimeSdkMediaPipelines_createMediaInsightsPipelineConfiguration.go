package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkMediaPipelines_createMediaInsightsPipelineConfigurationCmd = &cobra.Command{
	Use:   "create-media-insights-pipeline-configuration",
	Short: "A structure that contains the static configurations for a media insights pipeline.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkMediaPipelines_createMediaInsightsPipelineConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(chimeSdkMediaPipelines_createMediaInsightsPipelineConfigurationCmd).Standalone()

		chimeSdkMediaPipelines_createMediaInsightsPipelineConfigurationCmd.Flags().String("client-request-token", "", "The unique identifier for the media insights pipeline configuration request.")
		chimeSdkMediaPipelines_createMediaInsightsPipelineConfigurationCmd.Flags().String("elements", "", "The elements in the request, such as a processor for Amazon Transcribe or a sink for a Kinesis Data Stream.")
		chimeSdkMediaPipelines_createMediaInsightsPipelineConfigurationCmd.Flags().String("media-insights-pipeline-configuration-name", "", "The name of the media insights pipeline configuration.")
		chimeSdkMediaPipelines_createMediaInsightsPipelineConfigurationCmd.Flags().String("real-time-alert-configuration", "", "The configuration settings for the real-time alerts in a media insights pipeline configuration.")
		chimeSdkMediaPipelines_createMediaInsightsPipelineConfigurationCmd.Flags().String("resource-access-role-arn", "", "The ARN of the role used by the service to access Amazon Web Services resources, including `Transcribe` and `Transcribe Call Analytics`, on the caller’s behalf.")
		chimeSdkMediaPipelines_createMediaInsightsPipelineConfigurationCmd.Flags().String("tags", "", "The tags assigned to the media insights pipeline configuration.")
		chimeSdkMediaPipelines_createMediaInsightsPipelineConfigurationCmd.MarkFlagRequired("elements")
		chimeSdkMediaPipelines_createMediaInsightsPipelineConfigurationCmd.MarkFlagRequired("media-insights-pipeline-configuration-name")
		chimeSdkMediaPipelines_createMediaInsightsPipelineConfigurationCmd.MarkFlagRequired("resource-access-role-arn")
	})
	chimeSdkMediaPipelinesCmd.AddCommand(chimeSdkMediaPipelines_createMediaInsightsPipelineConfigurationCmd)
}
