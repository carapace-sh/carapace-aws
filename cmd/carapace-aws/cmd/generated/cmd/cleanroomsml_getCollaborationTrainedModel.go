package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cleanroomsml_getCollaborationTrainedModelCmd = &cobra.Command{
	Use:   "get-collaboration-trained-model",
	Short: "Returns information about a trained model in a collaboration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cleanroomsml_getCollaborationTrainedModelCmd).Standalone()

	cleanroomsml_getCollaborationTrainedModelCmd.Flags().String("collaboration-identifier", "", "The collaboration ID that contains the trained model that you want to return information about.")
	cleanroomsml_getCollaborationTrainedModelCmd.Flags().String("trained-model-arn", "", "The Amazon Resource Name (ARN) of the trained model that you want to return information about.")
	cleanroomsml_getCollaborationTrainedModelCmd.Flags().String("version-identifier", "", "The version identifier of the trained model to retrieve.")
	cleanroomsml_getCollaborationTrainedModelCmd.MarkFlagRequired("collaboration-identifier")
	cleanroomsml_getCollaborationTrainedModelCmd.MarkFlagRequired("trained-model-arn")
	cleanroomsmlCmd.AddCommand(cleanroomsml_getCollaborationTrainedModelCmd)
}
