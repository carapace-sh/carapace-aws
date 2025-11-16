package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_updatePipelineCmd = &cobra.Command{
	Use:   "update-pipeline",
	Short: "Updates a pipeline.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_updatePipelineCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sagemaker_updatePipelineCmd).Standalone()

		sagemaker_updatePipelineCmd.Flags().String("parallelism-configuration", "", "If specified, it applies to all executions of this pipeline by default.")
		sagemaker_updatePipelineCmd.Flags().String("pipeline-definition", "", "The JSON pipeline definition.")
		sagemaker_updatePipelineCmd.Flags().String("pipeline-definition-s3-location", "", "The location of the pipeline definition stored in Amazon S3.")
		sagemaker_updatePipelineCmd.Flags().String("pipeline-description", "", "The description of the pipeline.")
		sagemaker_updatePipelineCmd.Flags().String("pipeline-display-name", "", "The display name of the pipeline.")
		sagemaker_updatePipelineCmd.Flags().String("pipeline-name", "", "The name of the pipeline to update.")
		sagemaker_updatePipelineCmd.Flags().String("role-arn", "", "The Amazon Resource Name (ARN) that the pipeline uses to execute.")
		sagemaker_updatePipelineCmd.MarkFlagRequired("pipeline-name")
	})
	sagemakerCmd.AddCommand(sagemaker_updatePipelineCmd)
}
