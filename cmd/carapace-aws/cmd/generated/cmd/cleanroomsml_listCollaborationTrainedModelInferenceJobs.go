package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cleanroomsml_listCollaborationTrainedModelInferenceJobsCmd = &cobra.Command{
	Use:   "list-collaboration-trained-model-inference-jobs",
	Short: "Returns a list of trained model inference jobs in a specified collaboration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cleanroomsml_listCollaborationTrainedModelInferenceJobsCmd).Standalone()

	cleanroomsml_listCollaborationTrainedModelInferenceJobsCmd.Flags().String("collaboration-identifier", "", "The collaboration ID of the collaboration that contains the trained model inference jobs that you are interested in.")
	cleanroomsml_listCollaborationTrainedModelInferenceJobsCmd.Flags().String("max-results", "", "The maximum size of the results that is returned per call.")
	cleanroomsml_listCollaborationTrainedModelInferenceJobsCmd.Flags().String("next-token", "", "The token value retrieved from a previous call to access the next page of results.")
	cleanroomsml_listCollaborationTrainedModelInferenceJobsCmd.Flags().String("trained-model-arn", "", "The Amazon Resource Name (ARN) of the trained model that was used to create the trained model inference jobs that you are interested in.")
	cleanroomsml_listCollaborationTrainedModelInferenceJobsCmd.Flags().String("trained-model-version-identifier", "", "The version identifier of the trained model to filter inference jobs by.")
	cleanroomsml_listCollaborationTrainedModelInferenceJobsCmd.MarkFlagRequired("collaboration-identifier")
	cleanroomsmlCmd.AddCommand(cleanroomsml_listCollaborationTrainedModelInferenceJobsCmd)
}
