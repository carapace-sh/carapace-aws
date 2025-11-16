package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cleanroomsml_getTrainedModelCmd = &cobra.Command{
	Use:   "get-trained-model",
	Short: "Returns information about a trained model.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cleanroomsml_getTrainedModelCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cleanroomsml_getTrainedModelCmd).Standalone()

		cleanroomsml_getTrainedModelCmd.Flags().String("membership-identifier", "", "The membership ID of the member that created the trained model that you are interested in.")
		cleanroomsml_getTrainedModelCmd.Flags().String("trained-model-arn", "", "The Amazon Resource Name (ARN) of the trained model that you are interested in.")
		cleanroomsml_getTrainedModelCmd.Flags().String("version-identifier", "", "The version identifier of the trained model to retrieve.")
		cleanroomsml_getTrainedModelCmd.MarkFlagRequired("membership-identifier")
		cleanroomsml_getTrainedModelCmd.MarkFlagRequired("trained-model-arn")
	})
	cleanroomsmlCmd.AddCommand(cleanroomsml_getTrainedModelCmd)
}
