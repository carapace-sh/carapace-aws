package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkMediaPipelines_createMediaConcatenationPipelineCmd = &cobra.Command{
	Use:   "create-media-concatenation-pipeline",
	Short: "Creates a media concatenation pipeline.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkMediaPipelines_createMediaConcatenationPipelineCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(chimeSdkMediaPipelines_createMediaConcatenationPipelineCmd).Standalone()

		chimeSdkMediaPipelines_createMediaConcatenationPipelineCmd.Flags().String("client-request-token", "", "The unique identifier for the client request.")
		chimeSdkMediaPipelines_createMediaConcatenationPipelineCmd.Flags().String("sinks", "", "An object that specifies the data sinks for the media concatenation pipeline.")
		chimeSdkMediaPipelines_createMediaConcatenationPipelineCmd.Flags().String("sources", "", "An object that specifies the sources for the media concatenation pipeline.")
		chimeSdkMediaPipelines_createMediaConcatenationPipelineCmd.Flags().String("tags", "", "The tags associated with the media concatenation pipeline.")
		chimeSdkMediaPipelines_createMediaConcatenationPipelineCmd.MarkFlagRequired("sinks")
		chimeSdkMediaPipelines_createMediaConcatenationPipelineCmd.MarkFlagRequired("sources")
	})
	chimeSdkMediaPipelinesCmd.AddCommand(chimeSdkMediaPipelines_createMediaConcatenationPipelineCmd)
}
