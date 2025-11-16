package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_createOtaupdateCmd = &cobra.Command{
	Use:   "create-otaupdate",
	Short: "Creates an IoT OTA update on a target group of things or groups.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_createOtaupdateCmd).Standalone()

	iot_createOtaupdateCmd.Flags().String("additional-parameters", "", "A list of additional OTA update parameters, which are name-value pairs.")
	iot_createOtaupdateCmd.Flags().String("aws-job-abort-config", "", "The criteria that determine when and how a job abort takes place.")
	iot_createOtaupdateCmd.Flags().String("aws-job-executions-rollout-config", "", "Configuration for the rollout of OTA updates.")
	iot_createOtaupdateCmd.Flags().String("aws-job-presigned-url-config", "", "Configuration information for pre-signed URLs.")
	iot_createOtaupdateCmd.Flags().String("aws-job-timeout-config", "", "Specifies the amount of time each device has to finish its execution of the job.")
	iot_createOtaupdateCmd.Flags().String("description", "", "The description of the OTA update.")
	iot_createOtaupdateCmd.Flags().String("files", "", "The files to be streamed by the OTA update.")
	iot_createOtaupdateCmd.Flags().String("ota-update-id", "", "The ID of the OTA update to be created.")
	iot_createOtaupdateCmd.Flags().String("protocols", "", "The protocol used to transfer the OTA update image.")
	iot_createOtaupdateCmd.Flags().String("role-arn", "", "The IAM role that grants Amazon Web Services IoT Core access to the Amazon S3, IoT jobs and Amazon Web Services Code Signing resources to create an OTA update job.")
	iot_createOtaupdateCmd.Flags().String("tags", "", "Metadata which can be used to manage updates.")
	iot_createOtaupdateCmd.Flags().String("target-selection", "", "Specifies whether the update will continue to run (CONTINUOUS), or will be complete after all the things specified as targets have completed the update (SNAPSHOT).")
	iot_createOtaupdateCmd.Flags().String("targets", "", "The devices targeted to receive OTA updates.")
	iot_createOtaupdateCmd.MarkFlagRequired("files")
	iot_createOtaupdateCmd.MarkFlagRequired("ota-update-id")
	iot_createOtaupdateCmd.MarkFlagRequired("role-arn")
	iot_createOtaupdateCmd.MarkFlagRequired("targets")
	iotCmd.AddCommand(iot_createOtaupdateCmd)
}
