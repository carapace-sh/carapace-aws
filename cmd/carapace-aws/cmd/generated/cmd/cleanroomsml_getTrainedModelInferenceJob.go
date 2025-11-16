package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cleanroomsml_getTrainedModelInferenceJobCmd = &cobra.Command{
	Use:   "get-trained-model-inference-job",
	Short: "Returns information about a trained model inference job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cleanroomsml_getTrainedModelInferenceJobCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cleanroomsml_getTrainedModelInferenceJobCmd).Standalone()

		cleanroomsml_getTrainedModelInferenceJobCmd.Flags().String("membership-identifier", "", "Provides the membership ID of the membership that contains the trained model inference job that you are interested in.")
		cleanroomsml_getTrainedModelInferenceJobCmd.Flags().String("trained-model-inference-job-arn", "", "Provides the Amazon Resource Name (ARN) of the trained model inference job that you are interested in.")
		cleanroomsml_getTrainedModelInferenceJobCmd.MarkFlagRequired("membership-identifier")
		cleanroomsml_getTrainedModelInferenceJobCmd.MarkFlagRequired("trained-model-inference-job-arn")
	})
	cleanroomsmlCmd.AddCommand(cleanroomsml_getTrainedModelInferenceJobCmd)
}
