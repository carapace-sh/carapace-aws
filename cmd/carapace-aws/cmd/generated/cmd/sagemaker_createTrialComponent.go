package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_createTrialComponentCmd = &cobra.Command{
	Use:   "create-trial-component",
	Short: "Creates a *trial component*, which is a stage of a machine learning *trial*.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_createTrialComponentCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sagemaker_createTrialComponentCmd).Standalone()

		sagemaker_createTrialComponentCmd.Flags().String("display-name", "", "The name of the component as displayed.")
		sagemaker_createTrialComponentCmd.Flags().String("end-time", "", "When the component ended.")
		sagemaker_createTrialComponentCmd.Flags().String("input-artifacts", "", "The input artifacts for the component.")
		sagemaker_createTrialComponentCmd.Flags().String("metadata-properties", "", "")
		sagemaker_createTrialComponentCmd.Flags().String("output-artifacts", "", "The output artifacts for the component.")
		sagemaker_createTrialComponentCmd.Flags().String("parameters", "", "The hyperparameters for the component.")
		sagemaker_createTrialComponentCmd.Flags().String("start-time", "", "When the component started.")
		sagemaker_createTrialComponentCmd.Flags().String("status", "", "The status of the component.")
		sagemaker_createTrialComponentCmd.Flags().String("tags", "", "A list of tags to associate with the component.")
		sagemaker_createTrialComponentCmd.Flags().String("trial-component-name", "", "The name of the component.")
		sagemaker_createTrialComponentCmd.MarkFlagRequired("trial-component-name")
	})
	sagemakerCmd.AddCommand(sagemaker_createTrialComponentCmd)
}
