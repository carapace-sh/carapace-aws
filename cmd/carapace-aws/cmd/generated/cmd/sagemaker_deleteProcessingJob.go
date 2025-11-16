package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_deleteProcessingJobCmd = &cobra.Command{
	Use:   "delete-processing-job",
	Short: "Deletes a processing job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_deleteProcessingJobCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sagemaker_deleteProcessingJobCmd).Standalone()

		sagemaker_deleteProcessingJobCmd.Flags().String("processing-job-name", "", "The name of the processing job to delete.")
		sagemaker_deleteProcessingJobCmd.MarkFlagRequired("processing-job-name")
	})
	sagemakerCmd.AddCommand(sagemaker_deleteProcessingJobCmd)
}
