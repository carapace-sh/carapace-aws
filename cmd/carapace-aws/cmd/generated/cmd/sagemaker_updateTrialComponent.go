package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_updateTrialComponentCmd = &cobra.Command{
	Use:   "update-trial-component",
	Short: "Updates one or more properties of a trial component.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_updateTrialComponentCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sagemaker_updateTrialComponentCmd).Standalone()

		sagemaker_updateTrialComponentCmd.Flags().String("display-name", "", "The name of the component as displayed.")
		sagemaker_updateTrialComponentCmd.Flags().String("end-time", "", "When the component ended.")
		sagemaker_updateTrialComponentCmd.Flags().String("input-artifacts", "", "Replaces all of the component's input artifacts with the specified artifacts or adds new input artifacts.")
		sagemaker_updateTrialComponentCmd.Flags().String("input-artifacts-to-remove", "", "The input artifacts to remove from the component.")
		sagemaker_updateTrialComponentCmd.Flags().String("output-artifacts", "", "Replaces all of the component's output artifacts with the specified artifacts or adds new output artifacts.")
		sagemaker_updateTrialComponentCmd.Flags().String("output-artifacts-to-remove", "", "The output artifacts to remove from the component.")
		sagemaker_updateTrialComponentCmd.Flags().String("parameters", "", "Replaces all of the component's hyperparameters with the specified hyperparameters or add new hyperparameters.")
		sagemaker_updateTrialComponentCmd.Flags().String("parameters-to-remove", "", "The hyperparameters to remove from the component.")
		sagemaker_updateTrialComponentCmd.Flags().String("start-time", "", "When the component started.")
		sagemaker_updateTrialComponentCmd.Flags().String("status", "", "The new status of the component.")
		sagemaker_updateTrialComponentCmd.Flags().String("trial-component-name", "", "The name of the component to update.")
		sagemaker_updateTrialComponentCmd.MarkFlagRequired("trial-component-name")
	})
	sagemakerCmd.AddCommand(sagemaker_updateTrialComponentCmd)
}
