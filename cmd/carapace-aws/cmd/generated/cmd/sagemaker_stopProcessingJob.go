package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_stopProcessingJobCmd = &cobra.Command{
	Use:   "stop-processing-job",
	Short: "Stops a processing job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_stopProcessingJobCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sagemaker_stopProcessingJobCmd).Standalone()

		sagemaker_stopProcessingJobCmd.Flags().String("processing-job-name", "", "The name of the processing job to stop.")
		sagemaker_stopProcessingJobCmd.MarkFlagRequired("processing-job-name")
	})
	sagemakerCmd.AddCommand(sagemaker_stopProcessingJobCmd)
}
