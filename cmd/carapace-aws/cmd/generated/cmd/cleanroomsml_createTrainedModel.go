package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cleanroomsml_createTrainedModelCmd = &cobra.Command{
	Use:   "create-trained-model",
	Short: "Creates a trained model from an associated configured model algorithm using data from any member of the collaboration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cleanroomsml_createTrainedModelCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cleanroomsml_createTrainedModelCmd).Standalone()

		cleanroomsml_createTrainedModelCmd.Flags().String("configured-model-algorithm-association-arn", "", "The associated configured model algorithm used to train this model.")
		cleanroomsml_createTrainedModelCmd.Flags().String("data-channels", "", "Defines the data channels that are used as input for the trained model request.")
		cleanroomsml_createTrainedModelCmd.Flags().String("description", "", "The description of the trained model.")
		cleanroomsml_createTrainedModelCmd.Flags().String("environment", "", "The environment variables to set in the Docker container.")
		cleanroomsml_createTrainedModelCmd.Flags().String("hyperparameters", "", "Algorithm-specific parameters that influence the quality of the model.")
		cleanroomsml_createTrainedModelCmd.Flags().String("incremental-training-data-channels", "", "Specifies the incremental training data channels for the trained model.")
		cleanroomsml_createTrainedModelCmd.Flags().String("kms-key-arn", "", "The Amazon Resource Name (ARN) of the KMS key.")
		cleanroomsml_createTrainedModelCmd.Flags().String("membership-identifier", "", "The membership ID of the member that is creating the trained model.")
		cleanroomsml_createTrainedModelCmd.Flags().String("name", "", "The name of the trained model.")
		cleanroomsml_createTrainedModelCmd.Flags().String("resource-config", "", "Information about the EC2 resources that are used to train this model.")
		cleanroomsml_createTrainedModelCmd.Flags().String("stopping-condition", "", "The criteria that is used to stop model training.")
		cleanroomsml_createTrainedModelCmd.Flags().String("tags", "", "The optional metadata that you apply to the resource to help you categorize and organize them.")
		cleanroomsml_createTrainedModelCmd.Flags().String("training-input-mode", "", "The input mode for accessing the training data.")
		cleanroomsml_createTrainedModelCmd.MarkFlagRequired("configured-model-algorithm-association-arn")
		cleanroomsml_createTrainedModelCmd.MarkFlagRequired("data-channels")
		cleanroomsml_createTrainedModelCmd.MarkFlagRequired("membership-identifier")
		cleanroomsml_createTrainedModelCmd.MarkFlagRequired("name")
		cleanroomsml_createTrainedModelCmd.MarkFlagRequired("resource-config")
	})
	cleanroomsmlCmd.AddCommand(cleanroomsml_createTrainedModelCmd)
}
