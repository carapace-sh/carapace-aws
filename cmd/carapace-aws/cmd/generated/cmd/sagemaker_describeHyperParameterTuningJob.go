package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_describeHyperParameterTuningJobCmd = &cobra.Command{
	Use:   "describe-hyper-parameter-tuning-job",
	Short: "Returns a description of a hyperparameter tuning job, depending on the fields selected.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_describeHyperParameterTuningJobCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sagemaker_describeHyperParameterTuningJobCmd).Standalone()

		sagemaker_describeHyperParameterTuningJobCmd.Flags().String("hyper-parameter-tuning-job-name", "", "The name of the tuning job.")
		sagemaker_describeHyperParameterTuningJobCmd.MarkFlagRequired("hyper-parameter-tuning-job-name")
	})
	sagemakerCmd.AddCommand(sagemaker_describeHyperParameterTuningJobCmd)
}
