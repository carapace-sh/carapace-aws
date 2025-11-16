package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_describeInferenceRecommendationsJobCmd = &cobra.Command{
	Use:   "describe-inference-recommendations-job",
	Short: "Provides the results of the Inference Recommender job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_describeInferenceRecommendationsJobCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sagemaker_describeInferenceRecommendationsJobCmd).Standalone()

		sagemaker_describeInferenceRecommendationsJobCmd.Flags().String("job-name", "", "The name of the job.")
		sagemaker_describeInferenceRecommendationsJobCmd.MarkFlagRequired("job-name")
	})
	sagemakerCmd.AddCommand(sagemaker_describeInferenceRecommendationsJobCmd)
}
