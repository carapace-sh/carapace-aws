package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_deleteHyperParameterTuningJobCmd = &cobra.Command{
	Use:   "delete-hyper-parameter-tuning-job",
	Short: "Deletes a hyperparameter tuning job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_deleteHyperParameterTuningJobCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sagemaker_deleteHyperParameterTuningJobCmd).Standalone()

		sagemaker_deleteHyperParameterTuningJobCmd.Flags().String("hyper-parameter-tuning-job-name", "", "The name of the hyperparameter tuning job that you want to delete.")
		sagemaker_deleteHyperParameterTuningJobCmd.MarkFlagRequired("hyper-parameter-tuning-job-name")
	})
	sagemakerCmd.AddCommand(sagemaker_deleteHyperParameterTuningJobCmd)
}
