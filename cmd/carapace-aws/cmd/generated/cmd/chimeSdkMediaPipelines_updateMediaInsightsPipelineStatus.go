package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkMediaPipelines_updateMediaInsightsPipelineStatusCmd = &cobra.Command{
	Use:   "update-media-insights-pipeline-status",
	Short: "Updates the status of a media insights pipeline.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkMediaPipelines_updateMediaInsightsPipelineStatusCmd).Standalone()

	chimeSdkMediaPipelines_updateMediaInsightsPipelineStatusCmd.Flags().String("identifier", "", "The unique identifier of the resource to be updated.")
	chimeSdkMediaPipelines_updateMediaInsightsPipelineStatusCmd.Flags().String("update-status", "", "The requested status of the media insights pipeline.")
	chimeSdkMediaPipelines_updateMediaInsightsPipelineStatusCmd.MarkFlagRequired("identifier")
	chimeSdkMediaPipelines_updateMediaInsightsPipelineStatusCmd.MarkFlagRequired("update-status")
	chimeSdkMediaPipelinesCmd.AddCommand(chimeSdkMediaPipelines_updateMediaInsightsPipelineStatusCmd)
}
