package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cleanroomsml_getTrainingDatasetCmd = &cobra.Command{
	Use:   "get-training-dataset",
	Short: "Returns information about a training dataset.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cleanroomsml_getTrainingDatasetCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cleanroomsml_getTrainingDatasetCmd).Standalone()

		cleanroomsml_getTrainingDatasetCmd.Flags().String("training-dataset-arn", "", "The Amazon Resource Name (ARN) of the training dataset that you are interested in.")
		cleanroomsml_getTrainingDatasetCmd.MarkFlagRequired("training-dataset-arn")
	})
	cleanroomsmlCmd.AddCommand(cleanroomsml_getTrainingDatasetCmd)
}
