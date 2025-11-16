package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_updateJobCmd = &cobra.Command{
	Use:   "update-job",
	Short: "Updates supported fields of the specified job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_updateJobCmd).Standalone()

	iot_updateJobCmd.Flags().String("abort-config", "", "Allows you to create criteria to abort a job.")
	iot_updateJobCmd.Flags().String("description", "", "A short text description of the job.")
	iot_updateJobCmd.Flags().String("job-executions-retry-config", "", "Allows you to create the criteria to retry a job.")
	iot_updateJobCmd.Flags().String("job-executions-rollout-config", "", "Allows you to create a staged rollout of the job.")
	iot_updateJobCmd.Flags().String("job-id", "", "The ID of the job to be updated.")
	iot_updateJobCmd.Flags().String("namespace-id", "", "The namespace used to indicate that a job is a customer-managed job.")
	iot_updateJobCmd.Flags().String("presigned-url-config", "", "Configuration information for pre-signed S3 URLs.")
	iot_updateJobCmd.Flags().String("timeout-config", "", "Specifies the amount of time each device has to finish its execution of the job.")
	iot_updateJobCmd.MarkFlagRequired("job-id")
	iotCmd.AddCommand(iot_updateJobCmd)
}
