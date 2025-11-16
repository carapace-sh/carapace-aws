package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkMediaPipelines_getMediaInsightsPipelineConfigurationCmd = &cobra.Command{
	Use:   "get-media-insights-pipeline-configuration",
	Short: "Gets the configuration settings for a media insights pipeline.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkMediaPipelines_getMediaInsightsPipelineConfigurationCmd).Standalone()

	chimeSdkMediaPipelines_getMediaInsightsPipelineConfigurationCmd.Flags().String("identifier", "", "The unique identifier of the requested resource.")
	chimeSdkMediaPipelines_getMediaInsightsPipelineConfigurationCmd.MarkFlagRequired("identifier")
	chimeSdkMediaPipelinesCmd.AddCommand(chimeSdkMediaPipelines_getMediaInsightsPipelineConfigurationCmd)
}
