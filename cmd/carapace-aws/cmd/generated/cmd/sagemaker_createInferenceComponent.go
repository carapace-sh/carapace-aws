package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_createInferenceComponentCmd = &cobra.Command{
	Use:   "create-inference-component",
	Short: "Creates an inference component, which is a SageMaker AI hosting object that you can use to deploy a model to an endpoint.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_createInferenceComponentCmd).Standalone()

	sagemaker_createInferenceComponentCmd.Flags().String("endpoint-name", "", "The name of an existing endpoint where you host the inference component.")
	sagemaker_createInferenceComponentCmd.Flags().String("inference-component-name", "", "A unique name to assign to the inference component.")
	sagemaker_createInferenceComponentCmd.Flags().String("runtime-config", "", "Runtime settings for a model that is deployed with an inference component.")
	sagemaker_createInferenceComponentCmd.Flags().String("specification", "", "Details about the resources to deploy with this inference component, including the model, container, and compute resources.")
	sagemaker_createInferenceComponentCmd.Flags().String("tags", "", "A list of key-value pairs associated with the model.")
	sagemaker_createInferenceComponentCmd.Flags().String("variant-name", "", "The name of an existing production variant where you host the inference component.")
	sagemaker_createInferenceComponentCmd.MarkFlagRequired("endpoint-name")
	sagemaker_createInferenceComponentCmd.MarkFlagRequired("inference-component-name")
	sagemaker_createInferenceComponentCmd.MarkFlagRequired("specification")
	sagemakerCmd.AddCommand(sagemaker_createInferenceComponentCmd)
}
