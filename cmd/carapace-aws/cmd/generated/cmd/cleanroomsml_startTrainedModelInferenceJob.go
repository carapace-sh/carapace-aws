package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cleanroomsml_startTrainedModelInferenceJobCmd = &cobra.Command{
	Use:   "start-trained-model-inference-job",
	Short: "Defines the information necessary to begin a trained model inference job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cleanroomsml_startTrainedModelInferenceJobCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cleanroomsml_startTrainedModelInferenceJobCmd).Standalone()

		cleanroomsml_startTrainedModelInferenceJobCmd.Flags().String("configured-model-algorithm-association-arn", "", "The Amazon Resource Name (ARN) of the configured model algorithm association that is used for this trained model inference job.")
		cleanroomsml_startTrainedModelInferenceJobCmd.Flags().String("container-execution-parameters", "", "The execution parameters for the container.")
		cleanroomsml_startTrainedModelInferenceJobCmd.Flags().String("data-source", "", "Defines the data source that is used for the trained model inference job.")
		cleanroomsml_startTrainedModelInferenceJobCmd.Flags().String("description", "", "The description of the trained model inference job.")
		cleanroomsml_startTrainedModelInferenceJobCmd.Flags().String("environment", "", "The environment variables to set in the Docker container.")
		cleanroomsml_startTrainedModelInferenceJobCmd.Flags().String("kms-key-arn", "", "The Amazon Resource Name (ARN) of the KMS key.")
		cleanroomsml_startTrainedModelInferenceJobCmd.Flags().String("membership-identifier", "", "The membership ID of the membership that contains the trained model inference job.")
		cleanroomsml_startTrainedModelInferenceJobCmd.Flags().String("name", "", "The name of the trained model inference job.")
		cleanroomsml_startTrainedModelInferenceJobCmd.Flags().String("output-configuration", "", "Defines the output configuration information for the trained model inference job.")
		cleanroomsml_startTrainedModelInferenceJobCmd.Flags().String("resource-config", "", "Defines the resource configuration for the trained model inference job.")
		cleanroomsml_startTrainedModelInferenceJobCmd.Flags().String("tags", "", "The optional metadata that you apply to the resource to help you categorize and organize them.")
		cleanroomsml_startTrainedModelInferenceJobCmd.Flags().String("trained-model-arn", "", "The Amazon Resource Name (ARN) of the trained model that is used for this trained model inference job.")
		cleanroomsml_startTrainedModelInferenceJobCmd.Flags().String("trained-model-version-identifier", "", "The version identifier of the trained model to use for inference.")
		cleanroomsml_startTrainedModelInferenceJobCmd.MarkFlagRequired("data-source")
		cleanroomsml_startTrainedModelInferenceJobCmd.MarkFlagRequired("membership-identifier")
		cleanroomsml_startTrainedModelInferenceJobCmd.MarkFlagRequired("name")
		cleanroomsml_startTrainedModelInferenceJobCmd.MarkFlagRequired("output-configuration")
		cleanroomsml_startTrainedModelInferenceJobCmd.MarkFlagRequired("resource-config")
		cleanroomsml_startTrainedModelInferenceJobCmd.MarkFlagRequired("trained-model-arn")
	})
	cleanroomsmlCmd.AddCommand(cleanroomsml_startTrainedModelInferenceJobCmd)
}
