package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elastictranscoder_updatePipelineStatusCmd = &cobra.Command{
	Use:   "update-pipeline-status",
	Short: "The UpdatePipelineStatus operation pauses or reactivates a pipeline, so that the pipeline stops or restarts the processing of jobs.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elastictranscoder_updatePipelineStatusCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(elastictranscoder_updatePipelineStatusCmd).Standalone()

		elastictranscoder_updatePipelineStatusCmd.Flags().String("id", "", "The identifier of the pipeline to update.")
		elastictranscoder_updatePipelineStatusCmd.Flags().String("status", "", "The desired status of the pipeline:")
		elastictranscoder_updatePipelineStatusCmd.MarkFlagRequired("id")
		elastictranscoder_updatePipelineStatusCmd.MarkFlagRequired("status")
	})
	elastictranscoderCmd.AddCommand(elastictranscoder_updatePipelineStatusCmd)
}
