package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codepipeline_getPipelineStateCmd = &cobra.Command{
	Use:   "get-pipeline-state",
	Short: "Returns information about the state of a pipeline, including the stages and actions.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codepipeline_getPipelineStateCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codepipeline_getPipelineStateCmd).Standalone()

		codepipeline_getPipelineStateCmd.Flags().String("name", "", "The name of the pipeline about which you want to get information.")
		codepipeline_getPipelineStateCmd.MarkFlagRequired("name")
	})
	codepipelineCmd.AddCommand(codepipeline_getPipelineStateCmd)
}
