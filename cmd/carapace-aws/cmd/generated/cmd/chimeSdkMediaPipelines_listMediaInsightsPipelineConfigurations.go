package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkMediaPipelines_listMediaInsightsPipelineConfigurationsCmd = &cobra.Command{
	Use:   "list-media-insights-pipeline-configurations",
	Short: "Lists the available media insights pipeline configurations.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkMediaPipelines_listMediaInsightsPipelineConfigurationsCmd).Standalone()

	chimeSdkMediaPipelines_listMediaInsightsPipelineConfigurationsCmd.Flags().String("max-results", "", "The maximum number of results to return in a single call.")
	chimeSdkMediaPipelines_listMediaInsightsPipelineConfigurationsCmd.Flags().String("next-token", "", "The token used to return the next page of results.")
	chimeSdkMediaPipelinesCmd.AddCommand(chimeSdkMediaPipelines_listMediaInsightsPipelineConfigurationsCmd)
}
