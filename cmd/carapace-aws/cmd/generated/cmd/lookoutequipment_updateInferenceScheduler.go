package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lookoutequipment_updateInferenceSchedulerCmd = &cobra.Command{
	Use:   "update-inference-scheduler",
	Short: "Updates an inference scheduler.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lookoutequipment_updateInferenceSchedulerCmd).Standalone()

	lookoutequipment_updateInferenceSchedulerCmd.Flags().String("data-delay-offset-in-minutes", "", "A period of time (in minutes) by which inference on the data is delayed after the data starts.")
	lookoutequipment_updateInferenceSchedulerCmd.Flags().String("data-input-configuration", "", "Specifies information for the input data for the inference scheduler, including delimiter, format, and dataset location.")
	lookoutequipment_updateInferenceSchedulerCmd.Flags().String("data-output-configuration", "", "Specifies information for the output results from the inference scheduler, including the output S3 location.")
	lookoutequipment_updateInferenceSchedulerCmd.Flags().String("data-upload-frequency", "", "How often data is uploaded to the source S3 bucket for the input data.")
	lookoutequipment_updateInferenceSchedulerCmd.Flags().String("inference-scheduler-name", "", "The name of the inference scheduler to be updated.")
	lookoutequipment_updateInferenceSchedulerCmd.Flags().String("role-arn", "", "The Amazon Resource Name (ARN) of a role with permission to access the data source for the inference scheduler.")
	lookoutequipment_updateInferenceSchedulerCmd.MarkFlagRequired("inference-scheduler-name")
	lookoutequipmentCmd.AddCommand(lookoutequipment_updateInferenceSchedulerCmd)
}
