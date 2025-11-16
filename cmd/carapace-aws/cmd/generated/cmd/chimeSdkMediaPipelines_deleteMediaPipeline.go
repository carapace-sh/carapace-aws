package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkMediaPipelines_deleteMediaPipelineCmd = &cobra.Command{
	Use:   "delete-media-pipeline",
	Short: "Deletes the media pipeline.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkMediaPipelines_deleteMediaPipelineCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(chimeSdkMediaPipelines_deleteMediaPipelineCmd).Standalone()

		chimeSdkMediaPipelines_deleteMediaPipelineCmd.Flags().String("media-pipeline-id", "", "The ID of the media pipeline to delete.")
		chimeSdkMediaPipelines_deleteMediaPipelineCmd.MarkFlagRequired("media-pipeline-id")
	})
	chimeSdkMediaPipelinesCmd.AddCommand(chimeSdkMediaPipelines_deleteMediaPipelineCmd)
}
