package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cleanroomsml_startTrainedModelExportJobCmd = &cobra.Command{
	Use:   "start-trained-model-export-job",
	Short: "Provides the information necessary to start a trained model export job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cleanroomsml_startTrainedModelExportJobCmd).Standalone()

	cleanroomsml_startTrainedModelExportJobCmd.Flags().String("description", "", "The description of the trained model export job.")
	cleanroomsml_startTrainedModelExportJobCmd.Flags().String("membership-identifier", "", "The membership ID of the member that is receiving the exported trained model artifacts.")
	cleanroomsml_startTrainedModelExportJobCmd.Flags().String("name", "", "The name of the trained model export job.")
	cleanroomsml_startTrainedModelExportJobCmd.Flags().String("output-configuration", "", "The output configuration information for the trained model export job.")
	cleanroomsml_startTrainedModelExportJobCmd.Flags().String("trained-model-arn", "", "The Amazon Resource Name (ARN) of the trained model that you want to export.")
	cleanroomsml_startTrainedModelExportJobCmd.Flags().String("trained-model-version-identifier", "", "The version identifier of the trained model to export.")
	cleanroomsml_startTrainedModelExportJobCmd.MarkFlagRequired("membership-identifier")
	cleanroomsml_startTrainedModelExportJobCmd.MarkFlagRequired("name")
	cleanroomsml_startTrainedModelExportJobCmd.MarkFlagRequired("output-configuration")
	cleanroomsml_startTrainedModelExportJobCmd.MarkFlagRequired("trained-model-arn")
	cleanroomsmlCmd.AddCommand(cleanroomsml_startTrainedModelExportJobCmd)
}
