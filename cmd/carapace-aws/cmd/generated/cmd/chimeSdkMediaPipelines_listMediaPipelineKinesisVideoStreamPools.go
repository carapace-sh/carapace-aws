package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkMediaPipelines_listMediaPipelineKinesisVideoStreamPoolsCmd = &cobra.Command{
	Use:   "list-media-pipeline-kinesis-video-stream-pools",
	Short: "Lists the video stream pools in the media pipeline.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkMediaPipelines_listMediaPipelineKinesisVideoStreamPoolsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(chimeSdkMediaPipelines_listMediaPipelineKinesisVideoStreamPoolsCmd).Standalone()

		chimeSdkMediaPipelines_listMediaPipelineKinesisVideoStreamPoolsCmd.Flags().String("max-results", "", "The maximum number of results to return in a single call.")
		chimeSdkMediaPipelines_listMediaPipelineKinesisVideoStreamPoolsCmd.Flags().String("next-token", "", "The token used to return the next page of results.")
	})
	chimeSdkMediaPipelinesCmd.AddCommand(chimeSdkMediaPipelines_listMediaPipelineKinesisVideoStreamPoolsCmd)
}
