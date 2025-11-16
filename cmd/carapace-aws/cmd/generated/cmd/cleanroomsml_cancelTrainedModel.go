package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cleanroomsml_cancelTrainedModelCmd = &cobra.Command{
	Use:   "cancel-trained-model",
	Short: "Submits a request to cancel the trained model job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cleanroomsml_cancelTrainedModelCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cleanroomsml_cancelTrainedModelCmd).Standalone()

		cleanroomsml_cancelTrainedModelCmd.Flags().String("membership-identifier", "", "The membership ID of the trained model job that you want to cancel.")
		cleanroomsml_cancelTrainedModelCmd.Flags().String("trained-model-arn", "", "The Amazon Resource Name (ARN) of the trained model job that you want to cancel.")
		cleanroomsml_cancelTrainedModelCmd.Flags().String("version-identifier", "", "The version identifier of the trained model to cancel.")
		cleanroomsml_cancelTrainedModelCmd.MarkFlagRequired("membership-identifier")
		cleanroomsml_cancelTrainedModelCmd.MarkFlagRequired("trained-model-arn")
	})
	cleanroomsmlCmd.AddCommand(cleanroomsml_cancelTrainedModelCmd)
}
