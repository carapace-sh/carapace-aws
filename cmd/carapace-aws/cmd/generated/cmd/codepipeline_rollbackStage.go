package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codepipeline_rollbackStageCmd = &cobra.Command{
	Use:   "rollback-stage",
	Short: "Rolls back a stage execution.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codepipeline_rollbackStageCmd).Standalone()

	codepipeline_rollbackStageCmd.Flags().String("pipeline-name", "", "The name of the pipeline for which the stage will be rolled back.")
	codepipeline_rollbackStageCmd.Flags().String("stage-name", "", "The name of the stage in the pipeline to be rolled back.")
	codepipeline_rollbackStageCmd.Flags().String("target-pipeline-execution-id", "", "The pipeline execution ID for the stage to be rolled back to.")
	codepipeline_rollbackStageCmd.MarkFlagRequired("pipeline-name")
	codepipeline_rollbackStageCmd.MarkFlagRequired("stage-name")
	codepipeline_rollbackStageCmd.MarkFlagRequired("target-pipeline-execution-id")
	codepipelineCmd.AddCommand(codepipeline_rollbackStageCmd)
}
