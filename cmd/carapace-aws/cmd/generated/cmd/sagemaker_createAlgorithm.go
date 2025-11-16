package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_createAlgorithmCmd = &cobra.Command{
	Use:   "create-algorithm",
	Short: "Create a machine learning algorithm that you can use in SageMaker and list in the Amazon Web Services Marketplace.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_createAlgorithmCmd).Standalone()

	sagemaker_createAlgorithmCmd.Flags().String("algorithm-description", "", "A description of the algorithm.")
	sagemaker_createAlgorithmCmd.Flags().String("algorithm-name", "", "The name of the algorithm.")
	sagemaker_createAlgorithmCmd.Flags().String("certify-for-marketplace", "", "Whether to certify the algorithm so that it can be listed in Amazon Web Services Marketplace.")
	sagemaker_createAlgorithmCmd.Flags().String("inference-specification", "", "Specifies details about inference jobs that the algorithm runs, including the following:")
	sagemaker_createAlgorithmCmd.Flags().String("tags", "", "An array of key-value pairs.")
	sagemaker_createAlgorithmCmd.Flags().String("training-specification", "", "Specifies details about training jobs run by this algorithm, including the following:")
	sagemaker_createAlgorithmCmd.Flags().String("validation-specification", "", "Specifies configurations for one or more training jobs and that SageMaker runs to test the algorithm's training code and, optionally, one or more batch transform jobs that SageMaker runs to test the algorithm's inference code.")
	sagemaker_createAlgorithmCmd.MarkFlagRequired("algorithm-name")
	sagemaker_createAlgorithmCmd.MarkFlagRequired("training-specification")
	sagemakerCmd.AddCommand(sagemaker_createAlgorithmCmd)
}
