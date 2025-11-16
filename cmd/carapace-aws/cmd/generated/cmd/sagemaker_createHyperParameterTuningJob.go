package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_createHyperParameterTuningJobCmd = &cobra.Command{
	Use:   "create-hyper-parameter-tuning-job",
	Short: "Starts a hyperparameter tuning job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_createHyperParameterTuningJobCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sagemaker_createHyperParameterTuningJobCmd).Standalone()

		sagemaker_createHyperParameterTuningJobCmd.Flags().String("autotune", "", "Configures SageMaker Automatic model tuning (AMT) to automatically find optimal parameters for the following fields:")
		sagemaker_createHyperParameterTuningJobCmd.Flags().String("hyper-parameter-tuning-job-config", "", "The [HyperParameterTuningJobConfig](https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_HyperParameterTuningJobConfig.html) object that describes the tuning job, including the search strategy, the objective metric used to evaluate training jobs, ranges of parameters to search, and resource limits for the tuning job.")
		sagemaker_createHyperParameterTuningJobCmd.Flags().String("hyper-parameter-tuning-job-name", "", "The name of the tuning job.")
		sagemaker_createHyperParameterTuningJobCmd.Flags().String("tags", "", "An array of key-value pairs.")
		sagemaker_createHyperParameterTuningJobCmd.Flags().String("training-job-definition", "", "The [HyperParameterTrainingJobDefinition](https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_HyperParameterTrainingJobDefinition.html) object that describes the training jobs that this tuning job launches, including static hyperparameters, input data configuration, output data configuration, resource configuration, and stopping condition.")
		sagemaker_createHyperParameterTuningJobCmd.Flags().String("training-job-definitions", "", "A list of the [HyperParameterTrainingJobDefinition](https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_HyperParameterTrainingJobDefinition.html) objects launched for this tuning job.")
		sagemaker_createHyperParameterTuningJobCmd.Flags().String("warm-start-config", "", "Specifies the configuration for starting the hyperparameter tuning job using one or more previous tuning jobs as a starting point.")
		sagemaker_createHyperParameterTuningJobCmd.MarkFlagRequired("hyper-parameter-tuning-job-config")
		sagemaker_createHyperParameterTuningJobCmd.MarkFlagRequired("hyper-parameter-tuning-job-name")
	})
	sagemakerCmd.AddCommand(sagemaker_createHyperParameterTuningJobCmd)
}
