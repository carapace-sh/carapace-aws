package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cleanroomsml_listCollaborationTrainedModelsCmd = &cobra.Command{
	Use:   "list-collaboration-trained-models",
	Short: "Returns a list of the trained models in a collaboration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cleanroomsml_listCollaborationTrainedModelsCmd).Standalone()

	cleanroomsml_listCollaborationTrainedModelsCmd.Flags().String("collaboration-identifier", "", "The collaboration ID of the collaboration that contains the trained models you are interested in.")
	cleanroomsml_listCollaborationTrainedModelsCmd.Flags().String("max-results", "", "The maximum size of the results that is returned per call.")
	cleanroomsml_listCollaborationTrainedModelsCmd.Flags().String("next-token", "", "The token value retrieved from a previous call to access the next page of results.")
	cleanroomsml_listCollaborationTrainedModelsCmd.MarkFlagRequired("collaboration-identifier")
	cleanroomsmlCmd.AddCommand(cleanroomsml_listCollaborationTrainedModelsCmd)
}
