package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cleanroomsml_cancelTrainedModelInferenceJobCmd = &cobra.Command{
	Use:   "cancel-trained-model-inference-job",
	Short: "Submits a request to cancel a trained model inference job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cleanroomsml_cancelTrainedModelInferenceJobCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cleanroomsml_cancelTrainedModelInferenceJobCmd).Standalone()

		cleanroomsml_cancelTrainedModelInferenceJobCmd.Flags().String("membership-identifier", "", "The membership ID of the trained model inference job that you want to cancel.")
		cleanroomsml_cancelTrainedModelInferenceJobCmd.Flags().String("trained-model-inference-job-arn", "", "The Amazon Resource Name (ARN) of the trained model inference job that you want to cancel.")
		cleanroomsml_cancelTrainedModelInferenceJobCmd.MarkFlagRequired("membership-identifier")
		cleanroomsml_cancelTrainedModelInferenceJobCmd.MarkFlagRequired("trained-model-inference-job-arn")
	})
	cleanroomsmlCmd.AddCommand(cleanroomsml_cancelTrainedModelInferenceJobCmd)
}
