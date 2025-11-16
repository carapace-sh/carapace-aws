package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkMediaPipelines_createMediaStreamPipelineCmd = &cobra.Command{
	Use:   "create-media-stream-pipeline",
	Short: "Creates a streaming media pipeline.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkMediaPipelines_createMediaStreamPipelineCmd).Standalone()

	chimeSdkMediaPipelines_createMediaStreamPipelineCmd.Flags().String("client-request-token", "", "The token assigned to the client making the request.")
	chimeSdkMediaPipelines_createMediaStreamPipelineCmd.Flags().String("sinks", "", "The data sink for the media pipeline.")
	chimeSdkMediaPipelines_createMediaStreamPipelineCmd.Flags().String("sources", "", "The data sources for the media pipeline.")
	chimeSdkMediaPipelines_createMediaStreamPipelineCmd.Flags().String("tags", "", "The tags assigned to the media pipeline.")
	chimeSdkMediaPipelines_createMediaStreamPipelineCmd.MarkFlagRequired("sinks")
	chimeSdkMediaPipelines_createMediaStreamPipelineCmd.MarkFlagRequired("sources")
	chimeSdkMediaPipelinesCmd.AddCommand(chimeSdkMediaPipelines_createMediaStreamPipelineCmd)
}
