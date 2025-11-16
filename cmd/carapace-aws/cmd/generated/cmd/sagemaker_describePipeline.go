package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_describePipelineCmd = &cobra.Command{
	Use:   "describe-pipeline",
	Short: "Describes the details of a pipeline.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_describePipelineCmd).Standalone()

	sagemaker_describePipelineCmd.Flags().String("pipeline-name", "", "The name or Amazon Resource Name (ARN) of the pipeline to describe.")
	sagemaker_describePipelineCmd.Flags().String("pipeline-version-id", "", "The ID of the pipeline version to describe.")
	sagemaker_describePipelineCmd.MarkFlagRequired("pipeline-name")
	sagemakerCmd.AddCommand(sagemaker_describePipelineCmd)
}
