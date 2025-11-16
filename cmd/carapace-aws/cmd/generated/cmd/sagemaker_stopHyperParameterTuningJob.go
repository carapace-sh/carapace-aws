package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_stopHyperParameterTuningJobCmd = &cobra.Command{
	Use:   "stop-hyper-parameter-tuning-job",
	Short: "Stops a running hyperparameter tuning job and all running training jobs that the tuning job launched.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_stopHyperParameterTuningJobCmd).Standalone()

	sagemaker_stopHyperParameterTuningJobCmd.Flags().String("hyper-parameter-tuning-job-name", "", "The name of the tuning job to stop.")
	sagemaker_stopHyperParameterTuningJobCmd.MarkFlagRequired("hyper-parameter-tuning-job-name")
	sagemakerCmd.AddCommand(sagemaker_stopHyperParameterTuningJobCmd)
}
