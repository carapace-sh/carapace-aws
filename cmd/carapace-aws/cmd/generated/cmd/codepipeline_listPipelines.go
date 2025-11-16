package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codepipeline_listPipelinesCmd = &cobra.Command{
	Use:   "list-pipelines",
	Short: "Gets a summary of all of the pipelines associated with your account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codepipeline_listPipelinesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codepipeline_listPipelinesCmd).Standalone()

		codepipeline_listPipelinesCmd.Flags().String("max-results", "", "The maximum number of pipelines to return in a single call.")
		codepipeline_listPipelinesCmd.Flags().String("next-token", "", "An identifier that was returned from the previous list pipelines call.")
	})
	codepipelineCmd.AddCommand(codepipeline_listPipelinesCmd)
}
