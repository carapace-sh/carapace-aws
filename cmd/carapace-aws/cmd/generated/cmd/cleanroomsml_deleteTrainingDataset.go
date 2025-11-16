package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cleanroomsml_deleteTrainingDatasetCmd = &cobra.Command{
	Use:   "delete-training-dataset",
	Short: "Specifies a training dataset that you want to delete.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cleanroomsml_deleteTrainingDatasetCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cleanroomsml_deleteTrainingDatasetCmd).Standalone()

		cleanroomsml_deleteTrainingDatasetCmd.Flags().String("training-dataset-arn", "", "The Amazon Resource Name (ARN) of the training dataset that you want to delete.")
		cleanroomsml_deleteTrainingDatasetCmd.MarkFlagRequired("training-dataset-arn")
	})
	cleanroomsmlCmd.AddCommand(cleanroomsml_deleteTrainingDatasetCmd)
}
