package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_startMllabelingSetGenerationTaskRunCmd = &cobra.Command{
	Use:   "start-mllabeling-set-generation-task-run",
	Short: "Starts the active learning workflow for your machine learning transform to improve the transform's quality by generating label sets and adding labels.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_startMllabelingSetGenerationTaskRunCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(glue_startMllabelingSetGenerationTaskRunCmd).Standalone()

		glue_startMllabelingSetGenerationTaskRunCmd.Flags().String("output-s3-path", "", "The Amazon Simple Storage Service (Amazon S3) path where you generate the labeling set.")
		glue_startMllabelingSetGenerationTaskRunCmd.Flags().String("transform-id", "", "The unique identifier of the machine learning transform.")
		glue_startMllabelingSetGenerationTaskRunCmd.MarkFlagRequired("output-s3-path")
		glue_startMllabelingSetGenerationTaskRunCmd.MarkFlagRequired("transform-id")
	})
	glueCmd.AddCommand(glue_startMllabelingSetGenerationTaskRunCmd)
}
