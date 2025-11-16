package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cleanroomsml_deleteTrainedModelOutputCmd = &cobra.Command{
	Use:   "delete-trained-model-output",
	Short: "Deletes the model artifacts stored by the service.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cleanroomsml_deleteTrainedModelOutputCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cleanroomsml_deleteTrainedModelOutputCmd).Standalone()

		cleanroomsml_deleteTrainedModelOutputCmd.Flags().String("membership-identifier", "", "The membership ID of the member that is deleting the trained model output.")
		cleanroomsml_deleteTrainedModelOutputCmd.Flags().String("trained-model-arn", "", "The Amazon Resource Name (ARN) of the trained model whose output you want to delete.")
		cleanroomsml_deleteTrainedModelOutputCmd.Flags().String("version-identifier", "", "The version identifier of the trained model to delete.")
		cleanroomsml_deleteTrainedModelOutputCmd.MarkFlagRequired("membership-identifier")
		cleanroomsml_deleteTrainedModelOutputCmd.MarkFlagRequired("trained-model-arn")
	})
	cleanroomsmlCmd.AddCommand(cleanroomsml_deleteTrainedModelOutputCmd)
}
