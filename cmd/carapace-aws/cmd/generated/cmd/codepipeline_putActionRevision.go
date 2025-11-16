package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codepipeline_putActionRevisionCmd = &cobra.Command{
	Use:   "put-action-revision",
	Short: "Provides information to CodePipeline about new revisions to a source.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codepipeline_putActionRevisionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codepipeline_putActionRevisionCmd).Standalone()

		codepipeline_putActionRevisionCmd.Flags().String("action-name", "", "The name of the action that processes the revision.")
		codepipeline_putActionRevisionCmd.Flags().String("action-revision", "", "Represents information about the version (or revision) of an action.")
		codepipeline_putActionRevisionCmd.Flags().String("pipeline-name", "", "The name of the pipeline that starts processing the revision to the source.")
		codepipeline_putActionRevisionCmd.Flags().String("stage-name", "", "The name of the stage that contains the action that acts on the revision.")
		codepipeline_putActionRevisionCmd.MarkFlagRequired("action-name")
		codepipeline_putActionRevisionCmd.MarkFlagRequired("action-revision")
		codepipeline_putActionRevisionCmd.MarkFlagRequired("pipeline-name")
		codepipeline_putActionRevisionCmd.MarkFlagRequired("stage-name")
	})
	codepipelineCmd.AddCommand(codepipeline_putActionRevisionCmd)
}
