package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_createJobTemplateCmd = &cobra.Command{
	Use:   "create-job-template",
	Short: "Creates a job template.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_createJobTemplateCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iot_createJobTemplateCmd).Standalone()

		iot_createJobTemplateCmd.Flags().String("abort-config", "", "")
		iot_createJobTemplateCmd.Flags().String("description", "", "A description of the job document.")
		iot_createJobTemplateCmd.Flags().String("destination-package-versions", "", "The package version Amazon Resource Names (ARNs) that are installed on the device when the job successfully completes.")
		iot_createJobTemplateCmd.Flags().String("document", "", "The job document.")
		iot_createJobTemplateCmd.Flags().String("document-source", "", "An S3 link, or S3 object URL, to the job document.")
		iot_createJobTemplateCmd.Flags().String("job-arn", "", "The ARN of the job to use as the basis for the job template.")
		iot_createJobTemplateCmd.Flags().String("job-executions-retry-config", "", "Allows you to create the criteria to retry a job.")
		iot_createJobTemplateCmd.Flags().String("job-executions-rollout-config", "", "")
		iot_createJobTemplateCmd.Flags().String("job-template-id", "", "A unique identifier for the job template.")
		iot_createJobTemplateCmd.Flags().String("maintenance-windows", "", "Allows you to configure an optional maintenance window for the rollout of a job document to all devices in the target group for a job.")
		iot_createJobTemplateCmd.Flags().String("presigned-url-config", "", "")
		iot_createJobTemplateCmd.Flags().String("tags", "", "Metadata that can be used to manage the job template.")
		iot_createJobTemplateCmd.Flags().String("timeout-config", "", "")
		iot_createJobTemplateCmd.MarkFlagRequired("description")
		iot_createJobTemplateCmd.MarkFlagRequired("job-template-id")
	})
	iotCmd.AddCommand(iot_createJobTemplateCmd)
}
