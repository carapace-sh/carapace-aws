package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_createJobCmd = &cobra.Command{
	Use:   "create-job",
	Short: "Creates a job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_createJobCmd).Standalone()

	iot_createJobCmd.Flags().String("abort-config", "", "Allows you to create the criteria to abort a job.")
	iot_createJobCmd.Flags().String("description", "", "A short text description of the job.")
	iot_createJobCmd.Flags().String("destination-package-versions", "", "The package version Amazon Resource Names (ARNs) that are installed on the device when the job successfully completes.")
	iot_createJobCmd.Flags().String("document", "", "The job document.")
	iot_createJobCmd.Flags().String("document-parameters", "", "Parameters of an Amazon Web Services managed template that you can specify to create the job document.")
	iot_createJobCmd.Flags().String("document-source", "", "An S3 link, or S3 object URL, to the job document.")
	iot_createJobCmd.Flags().String("job-executions-retry-config", "", "Allows you to create the criteria to retry a job.")
	iot_createJobCmd.Flags().String("job-executions-rollout-config", "", "Allows you to create a staged rollout of the job.")
	iot_createJobCmd.Flags().String("job-id", "", "A job identifier which must be unique for your account.")
	iot_createJobCmd.Flags().String("job-template-arn", "", "The ARN of the job template used to create the job.")
	iot_createJobCmd.Flags().String("namespace-id", "", "The namespace used to indicate that a job is a customer-managed job.")
	iot_createJobCmd.Flags().String("presigned-url-config", "", "Configuration information for pre-signed S3 URLs.")
	iot_createJobCmd.Flags().String("scheduling-config", "", "The configuration that allows you to schedule a job for a future date and time in addition to specifying the end behavior for each job execution.")
	iot_createJobCmd.Flags().String("tags", "", "Metadata which can be used to manage the job.")
	iot_createJobCmd.Flags().String("target-selection", "", "Specifies whether the job will continue to run (CONTINUOUS), or will be complete after all those things specified as targets have completed the job (SNAPSHOT).")
	iot_createJobCmd.Flags().String("targets", "", "A list of things and thing groups to which the job should be sent.")
	iot_createJobCmd.Flags().String("timeout-config", "", "Specifies the amount of time each device has to finish its execution of the job.")
	iot_createJobCmd.MarkFlagRequired("job-id")
	iot_createJobCmd.MarkFlagRequired("targets")
	iotCmd.AddCommand(iot_createJobCmd)
}
