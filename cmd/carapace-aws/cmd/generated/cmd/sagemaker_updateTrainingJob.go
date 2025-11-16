package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_updateTrainingJobCmd = &cobra.Command{
	Use:   "update-training-job",
	Short: "Update a model training job to request a new Debugger profiling configuration or to change warm pool retention length.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_updateTrainingJobCmd).Standalone()

	sagemaker_updateTrainingJobCmd.Flags().String("profiler-config", "", "Configuration information for Amazon SageMaker Debugger system monitoring, framework profiling, and storage paths.")
	sagemaker_updateTrainingJobCmd.Flags().String("profiler-rule-configurations", "", "Configuration information for Amazon SageMaker Debugger rules for profiling system and framework metrics.")
	sagemaker_updateTrainingJobCmd.Flags().String("remote-debug-config", "", "Configuration for remote debugging while the training job is running.")
	sagemaker_updateTrainingJobCmd.Flags().String("resource-config", "", "The training job `ResourceConfig` to update warm pool retention length.")
	sagemaker_updateTrainingJobCmd.Flags().String("training-job-name", "", "The name of a training job to update the Debugger profiling configuration.")
	sagemaker_updateTrainingJobCmd.MarkFlagRequired("training-job-name")
	sagemakerCmd.AddCommand(sagemaker_updateTrainingJobCmd)
}
