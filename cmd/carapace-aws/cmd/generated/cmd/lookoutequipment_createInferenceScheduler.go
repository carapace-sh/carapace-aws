package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lookoutequipment_createInferenceSchedulerCmd = &cobra.Command{
	Use:   "create-inference-scheduler",
	Short: "Creates a scheduled inference.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lookoutequipment_createInferenceSchedulerCmd).Standalone()

	lookoutequipment_createInferenceSchedulerCmd.Flags().String("client-token", "", "A unique identifier for the request.")
	lookoutequipment_createInferenceSchedulerCmd.Flags().String("data-delay-offset-in-minutes", "", "The interval (in minutes) of planned delay at the start of each inference segment.")
	lookoutequipment_createInferenceSchedulerCmd.Flags().String("data-input-configuration", "", "Specifies configuration information for the input data for the inference scheduler, including delimiter, format, and dataset location.")
	lookoutequipment_createInferenceSchedulerCmd.Flags().String("data-output-configuration", "", "Specifies configuration information for the output results for the inference scheduler, including the S3 location for the output.")
	lookoutequipment_createInferenceSchedulerCmd.Flags().String("data-upload-frequency", "", "How often data is uploaded to the source Amazon S3 bucket for the input data.")
	lookoutequipment_createInferenceSchedulerCmd.Flags().String("inference-scheduler-name", "", "The name of the inference scheduler being created.")
	lookoutequipment_createInferenceSchedulerCmd.Flags().String("model-name", "", "The name of the previously trained machine learning model being used to create the inference scheduler.")
	lookoutequipment_createInferenceSchedulerCmd.Flags().String("role-arn", "", "The Amazon Resource Name (ARN) of a role with permission to access the data source being used for the inference.")
	lookoutequipment_createInferenceSchedulerCmd.Flags().String("server-side-kms-key-id", "", "Provides the identifier of the KMS key used to encrypt inference scheduler data by Amazon Lookout for Equipment.")
	lookoutequipment_createInferenceSchedulerCmd.Flags().String("tags", "", "Any tags associated with the inference scheduler.")
	lookoutequipment_createInferenceSchedulerCmd.MarkFlagRequired("client-token")
	lookoutequipment_createInferenceSchedulerCmd.MarkFlagRequired("data-input-configuration")
	lookoutequipment_createInferenceSchedulerCmd.MarkFlagRequired("data-output-configuration")
	lookoutequipment_createInferenceSchedulerCmd.MarkFlagRequired("data-upload-frequency")
	lookoutequipment_createInferenceSchedulerCmd.MarkFlagRequired("inference-scheduler-name")
	lookoutequipment_createInferenceSchedulerCmd.MarkFlagRequired("model-name")
	lookoutequipment_createInferenceSchedulerCmd.MarkFlagRequired("role-arn")
	lookoutequipmentCmd.AddCommand(lookoutequipment_createInferenceSchedulerCmd)
}
