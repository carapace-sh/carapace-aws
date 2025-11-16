package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_describeProcessingJobCmd = &cobra.Command{
	Use:   "describe-processing-job",
	Short: "Returns a description of a processing job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_describeProcessingJobCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sagemaker_describeProcessingJobCmd).Standalone()

		sagemaker_describeProcessingJobCmd.Flags().String("processing-job-name", "", "The name of the processing job.")
		sagemaker_describeProcessingJobCmd.MarkFlagRequired("processing-job-name")
	})
	sagemakerCmd.AddCommand(sagemaker_describeProcessingJobCmd)
}
