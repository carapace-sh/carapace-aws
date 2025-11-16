package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkMediaPipelines_deleteMediaInsightsPipelineConfigurationCmd = &cobra.Command{
	Use:   "delete-media-insights-pipeline-configuration",
	Short: "Deletes the specified configuration settings.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkMediaPipelines_deleteMediaInsightsPipelineConfigurationCmd).Standalone()

	chimeSdkMediaPipelines_deleteMediaInsightsPipelineConfigurationCmd.Flags().String("identifier", "", "The unique identifier of the resource to be deleted.")
	chimeSdkMediaPipelines_deleteMediaInsightsPipelineConfigurationCmd.MarkFlagRequired("identifier")
	chimeSdkMediaPipelinesCmd.AddCommand(chimeSdkMediaPipelines_deleteMediaInsightsPipelineConfigurationCmd)
}
