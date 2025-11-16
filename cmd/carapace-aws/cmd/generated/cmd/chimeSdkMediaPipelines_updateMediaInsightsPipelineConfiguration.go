package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkMediaPipelines_updateMediaInsightsPipelineConfigurationCmd = &cobra.Command{
	Use:   "update-media-insights-pipeline-configuration",
	Short: "Updates the media insights pipeline's configuration settings.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkMediaPipelines_updateMediaInsightsPipelineConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(chimeSdkMediaPipelines_updateMediaInsightsPipelineConfigurationCmd).Standalone()

		chimeSdkMediaPipelines_updateMediaInsightsPipelineConfigurationCmd.Flags().String("elements", "", "The elements in the request, such as a processor for Amazon Transcribe or a sink for a Kinesis Data Stream..")
		chimeSdkMediaPipelines_updateMediaInsightsPipelineConfigurationCmd.Flags().String("identifier", "", "The unique identifier for the resource to be updated.")
		chimeSdkMediaPipelines_updateMediaInsightsPipelineConfigurationCmd.Flags().String("real-time-alert-configuration", "", "The configuration settings for real-time alerts for the media insights pipeline.")
		chimeSdkMediaPipelines_updateMediaInsightsPipelineConfigurationCmd.Flags().String("resource-access-role-arn", "", "The ARN of the role used by the service to access Amazon Web Services resources.")
		chimeSdkMediaPipelines_updateMediaInsightsPipelineConfigurationCmd.MarkFlagRequired("elements")
		chimeSdkMediaPipelines_updateMediaInsightsPipelineConfigurationCmd.MarkFlagRequired("identifier")
		chimeSdkMediaPipelines_updateMediaInsightsPipelineConfigurationCmd.MarkFlagRequired("resource-access-role-arn")
	})
	chimeSdkMediaPipelinesCmd.AddCommand(chimeSdkMediaPipelines_updateMediaInsightsPipelineConfigurationCmd)
}
