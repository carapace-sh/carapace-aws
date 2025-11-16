package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_stopInferenceRecommendationsJobCmd = &cobra.Command{
	Use:   "stop-inference-recommendations-job",
	Short: "Stops an Inference Recommender job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_stopInferenceRecommendationsJobCmd).Standalone()

	sagemaker_stopInferenceRecommendationsJobCmd.Flags().String("job-name", "", "The name of the job you want to stop.")
	sagemaker_stopInferenceRecommendationsJobCmd.MarkFlagRequired("job-name")
	sagemakerCmd.AddCommand(sagemaker_stopInferenceRecommendationsJobCmd)
}
