package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkMediaPipelines_createMediaLiveConnectorPipelineCmd = &cobra.Command{
	Use:   "create-media-live-connector-pipeline",
	Short: "Creates a media live connector pipeline in an Amazon Chime SDK meeting.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkMediaPipelines_createMediaLiveConnectorPipelineCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(chimeSdkMediaPipelines_createMediaLiveConnectorPipelineCmd).Standalone()

		chimeSdkMediaPipelines_createMediaLiveConnectorPipelineCmd.Flags().String("client-request-token", "", "The token assigned to the client making the request.")
		chimeSdkMediaPipelines_createMediaLiveConnectorPipelineCmd.Flags().String("sinks", "", "The media live connector pipeline's data sinks.")
		chimeSdkMediaPipelines_createMediaLiveConnectorPipelineCmd.Flags().String("sources", "", "The media live connector pipeline's data sources.")
		chimeSdkMediaPipelines_createMediaLiveConnectorPipelineCmd.Flags().String("tags", "", "The tags associated with the media live connector pipeline.")
		chimeSdkMediaPipelines_createMediaLiveConnectorPipelineCmd.MarkFlagRequired("sinks")
		chimeSdkMediaPipelines_createMediaLiveConnectorPipelineCmd.MarkFlagRequired("sources")
	})
	chimeSdkMediaPipelinesCmd.AddCommand(chimeSdkMediaPipelines_createMediaLiveConnectorPipelineCmd)
}
