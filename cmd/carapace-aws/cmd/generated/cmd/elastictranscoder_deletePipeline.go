package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elastictranscoder_deletePipelineCmd = &cobra.Command{
	Use:   "delete-pipeline",
	Short: "The DeletePipeline operation removes a pipeline.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elastictranscoder_deletePipelineCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(elastictranscoder_deletePipelineCmd).Standalone()

		elastictranscoder_deletePipelineCmd.Flags().String("id", "", "The identifier of the pipeline that you want to delete.")
		elastictranscoder_deletePipelineCmd.MarkFlagRequired("id")
	})
	elastictranscoderCmd.AddCommand(elastictranscoder_deletePipelineCmd)
}
