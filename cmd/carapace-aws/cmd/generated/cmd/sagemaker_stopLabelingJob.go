package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_stopLabelingJobCmd = &cobra.Command{
	Use:   "stop-labeling-job",
	Short: "Stops a running labeling job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_stopLabelingJobCmd).Standalone()

	sagemaker_stopLabelingJobCmd.Flags().String("labeling-job-name", "", "The name of the labeling job to stop.")
	sagemaker_stopLabelingJobCmd.MarkFlagRequired("labeling-job-name")
	sagemakerCmd.AddCommand(sagemaker_stopLabelingJobCmd)
}
