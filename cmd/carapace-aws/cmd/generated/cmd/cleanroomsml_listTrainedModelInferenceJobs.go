package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cleanroomsml_listTrainedModelInferenceJobsCmd = &cobra.Command{
	Use:   "list-trained-model-inference-jobs",
	Short: "Returns a list of trained model inference jobs that match the request parameters.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cleanroomsml_listTrainedModelInferenceJobsCmd).Standalone()

	cleanroomsml_listTrainedModelInferenceJobsCmd.Flags().String("max-results", "", "The maximum size of the results that is returned per call.")
	cleanroomsml_listTrainedModelInferenceJobsCmd.Flags().String("membership-identifier", "", "The membership")
	cleanroomsml_listTrainedModelInferenceJobsCmd.Flags().String("next-token", "", "The token value retrieved from a previous call to access the next page of results.")
	cleanroomsml_listTrainedModelInferenceJobsCmd.Flags().String("trained-model-arn", "", "The Amazon Resource Name (ARN) of a trained model that was used to create the trained model inference jobs that you are interested in.")
	cleanroomsml_listTrainedModelInferenceJobsCmd.Flags().String("trained-model-version-identifier", "", "The version identifier of the trained model to filter inference jobs by.")
	cleanroomsml_listTrainedModelInferenceJobsCmd.MarkFlagRequired("membership-identifier")
	cleanroomsmlCmd.AddCommand(cleanroomsml_listTrainedModelInferenceJobsCmd)
}
