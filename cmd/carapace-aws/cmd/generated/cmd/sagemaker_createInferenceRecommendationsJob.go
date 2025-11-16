package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_createInferenceRecommendationsJobCmd = &cobra.Command{
	Use:   "create-inference-recommendations-job",
	Short: "Starts a recommendation job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_createInferenceRecommendationsJobCmd).Standalone()

	sagemaker_createInferenceRecommendationsJobCmd.Flags().String("input-config", "", "Provides information about the versioned model package Amazon Resource Name (ARN), the traffic pattern, and endpoint configurations.")
	sagemaker_createInferenceRecommendationsJobCmd.Flags().String("job-description", "", "Description of the recommendation job.")
	sagemaker_createInferenceRecommendationsJobCmd.Flags().String("job-name", "", "A name for the recommendation job.")
	sagemaker_createInferenceRecommendationsJobCmd.Flags().String("job-type", "", "Defines the type of recommendation job.")
	sagemaker_createInferenceRecommendationsJobCmd.Flags().String("output-config", "", "Provides information about the output artifacts and the KMS key to use for Amazon S3 server-side encryption.")
	sagemaker_createInferenceRecommendationsJobCmd.Flags().String("role-arn", "", "The Amazon Resource Name (ARN) of an IAM role that enables Amazon SageMaker to perform tasks on your behalf.")
	sagemaker_createInferenceRecommendationsJobCmd.Flags().String("stopping-conditions", "", "A set of conditions for stopping a recommendation job.")
	sagemaker_createInferenceRecommendationsJobCmd.Flags().String("tags", "", "The metadata that you apply to Amazon Web Services resources to help you categorize and organize them.")
	sagemaker_createInferenceRecommendationsJobCmd.MarkFlagRequired("input-config")
	sagemaker_createInferenceRecommendationsJobCmd.MarkFlagRequired("job-name")
	sagemaker_createInferenceRecommendationsJobCmd.MarkFlagRequired("job-type")
	sagemaker_createInferenceRecommendationsJobCmd.MarkFlagRequired("role-arn")
	sagemakerCmd.AddCommand(sagemaker_createInferenceRecommendationsJobCmd)
}
