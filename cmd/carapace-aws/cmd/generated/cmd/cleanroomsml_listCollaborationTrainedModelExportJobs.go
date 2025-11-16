package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cleanroomsml_listCollaborationTrainedModelExportJobsCmd = &cobra.Command{
	Use:   "list-collaboration-trained-model-export-jobs",
	Short: "Returns a list of the export jobs for a trained model in a collaboration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cleanroomsml_listCollaborationTrainedModelExportJobsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cleanroomsml_listCollaborationTrainedModelExportJobsCmd).Standalone()

		cleanroomsml_listCollaborationTrainedModelExportJobsCmd.Flags().String("collaboration-identifier", "", "The collaboration ID of the collaboration that contains the trained model export jobs that you are interested in.")
		cleanroomsml_listCollaborationTrainedModelExportJobsCmd.Flags().String("max-results", "", "The maximum size of the results that is returned per call.")
		cleanroomsml_listCollaborationTrainedModelExportJobsCmd.Flags().String("next-token", "", "The token value retrieved from a previous call to access the next page of results.")
		cleanroomsml_listCollaborationTrainedModelExportJobsCmd.Flags().String("trained-model-arn", "", "The Amazon Resource Name (ARN) of the trained model that was used to create the export jobs that you are interested in.")
		cleanroomsml_listCollaborationTrainedModelExportJobsCmd.Flags().String("trained-model-version-identifier", "", "The version identifier of the trained model to filter export jobs by.")
		cleanroomsml_listCollaborationTrainedModelExportJobsCmd.MarkFlagRequired("collaboration-identifier")
		cleanroomsml_listCollaborationTrainedModelExportJobsCmd.MarkFlagRequired("trained-model-arn")
	})
	cleanroomsmlCmd.AddCommand(cleanroomsml_listCollaborationTrainedModelExportJobsCmd)
}
