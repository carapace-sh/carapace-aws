package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lookoutequipment_createModelCmd = &cobra.Command{
	Use:   "create-model",
	Short: "Creates a machine learning model for data inference.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lookoutequipment_createModelCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lookoutequipment_createModelCmd).Standalone()

		lookoutequipment_createModelCmd.Flags().String("client-token", "", "A unique identifier for the request.")
		lookoutequipment_createModelCmd.Flags().String("data-pre-processing-configuration", "", "The configuration is the `TargetSamplingRate`, which is the sampling rate of the data after post processing by Amazon Lookout for Equipment.")
		lookoutequipment_createModelCmd.Flags().String("dataset-name", "", "The name of the dataset for the machine learning model being created.")
		lookoutequipment_createModelCmd.Flags().String("dataset-schema", "", "The data schema for the machine learning model being created.")
		lookoutequipment_createModelCmd.Flags().String("evaluation-data-end-time", "", "Indicates the time reference in the dataset that should be used to end the subset of evaluation data for the machine learning model.")
		lookoutequipment_createModelCmd.Flags().String("evaluation-data-start-time", "", "Indicates the time reference in the dataset that should be used to begin the subset of evaluation data for the machine learning model.")
		lookoutequipment_createModelCmd.Flags().String("labels-input-configuration", "", "The input configuration for the labels being used for the machine learning model that's being created.")
		lookoutequipment_createModelCmd.Flags().String("model-diagnostics-output-configuration", "", "The Amazon S3 location where you want Amazon Lookout for Equipment to save the pointwise model diagnostics.")
		lookoutequipment_createModelCmd.Flags().String("model-name", "", "The name for the machine learning model to be created.")
		lookoutequipment_createModelCmd.Flags().String("off-condition", "", "Indicates that the asset associated with this sensor has been shut off.")
		lookoutequipment_createModelCmd.Flags().String("role-arn", "", "The Amazon Resource Name (ARN) of a role with permission to access the data source being used to create the machine learning model.")
		lookoutequipment_createModelCmd.Flags().String("server-side-kms-key-id", "", "Provides the identifier of the KMS key used to encrypt model data by Amazon Lookout for Equipment.")
		lookoutequipment_createModelCmd.Flags().String("tags", "", "Any tags associated with the machine learning model being created.")
		lookoutequipment_createModelCmd.Flags().String("training-data-end-time", "", "Indicates the time reference in the dataset that should be used to end the subset of training data for the machine learning model.")
		lookoutequipment_createModelCmd.Flags().String("training-data-start-time", "", "Indicates the time reference in the dataset that should be used to begin the subset of training data for the machine learning model.")
		lookoutequipment_createModelCmd.MarkFlagRequired("client-token")
		lookoutequipment_createModelCmd.MarkFlagRequired("dataset-name")
		lookoutequipment_createModelCmd.MarkFlagRequired("model-name")
	})
	lookoutequipmentCmd.AddCommand(lookoutequipment_createModelCmd)
}
