package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elastictranscoder_readPipelineCmd = &cobra.Command{
	Use:   "read-pipeline",
	Short: "The ReadPipeline operation gets detailed information about a pipeline.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elastictranscoder_readPipelineCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(elastictranscoder_readPipelineCmd).Standalone()

		elastictranscoder_readPipelineCmd.Flags().String("id", "", "The identifier of the pipeline to read.")
		elastictranscoder_readPipelineCmd.MarkFlagRequired("id")
	})
	elastictranscoderCmd.AddCommand(elastictranscoder_readPipelineCmd)
}
