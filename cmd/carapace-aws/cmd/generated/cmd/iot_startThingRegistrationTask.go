package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_startThingRegistrationTaskCmd = &cobra.Command{
	Use:   "start-thing-registration-task",
	Short: "Creates a bulk thing provisioning task.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_startThingRegistrationTaskCmd).Standalone()

	iot_startThingRegistrationTaskCmd.Flags().String("input-file-bucket", "", "The S3 bucket that contains the input file.")
	iot_startThingRegistrationTaskCmd.Flags().String("input-file-key", "", "The name of input file within the S3 bucket.")
	iot_startThingRegistrationTaskCmd.Flags().String("role-arn", "", "The IAM role ARN that grants permission the input file.")
	iot_startThingRegistrationTaskCmd.Flags().String("template-body", "", "The provisioning template.")
	iot_startThingRegistrationTaskCmd.MarkFlagRequired("input-file-bucket")
	iot_startThingRegistrationTaskCmd.MarkFlagRequired("input-file-key")
	iot_startThingRegistrationTaskCmd.MarkFlagRequired("role-arn")
	iot_startThingRegistrationTaskCmd.MarkFlagRequired("template-body")
	iotCmd.AddCommand(iot_startThingRegistrationTaskCmd)
}
