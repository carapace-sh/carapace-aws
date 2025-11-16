package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssm_registerTaskWithMaintenanceWindowCmd = &cobra.Command{
	Use:   "register-task-with-maintenance-window",
	Short: "Adds a new task to a maintenance window.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssm_registerTaskWithMaintenanceWindowCmd).Standalone()

	ssm_registerTaskWithMaintenanceWindowCmd.Flags().String("alarm-configuration", "", "The CloudWatch alarm you want to apply to your maintenance window task.")
	ssm_registerTaskWithMaintenanceWindowCmd.Flags().String("client-token", "", "User-provided idempotency token.")
	ssm_registerTaskWithMaintenanceWindowCmd.Flags().String("cutoff-behavior", "", "Indicates whether tasks should continue to run after the cutoff time specified in the maintenance windows is reached.")
	ssm_registerTaskWithMaintenanceWindowCmd.Flags().String("description", "", "An optional description for the task.")
	ssm_registerTaskWithMaintenanceWindowCmd.Flags().String("logging-info", "", "A structure containing information about an Amazon Simple Storage Service (Amazon S3) bucket to write managed node-level logs to.")
	ssm_registerTaskWithMaintenanceWindowCmd.Flags().String("max-concurrency", "", "The maximum number of targets this task can be run for, in parallel.")
	ssm_registerTaskWithMaintenanceWindowCmd.Flags().String("max-errors", "", "The maximum number of errors allowed before this task stops being scheduled.")
	ssm_registerTaskWithMaintenanceWindowCmd.Flags().String("name", "", "An optional name for the task.")
	ssm_registerTaskWithMaintenanceWindowCmd.Flags().String("priority", "", "The priority of the task in the maintenance window, the lower the number the higher the priority.")
	ssm_registerTaskWithMaintenanceWindowCmd.Flags().String("service-role-arn", "", "The Amazon Resource Name (ARN) of the IAM service role for Amazon Web Services Systems Manager to assume when running a maintenance window task.")
	ssm_registerTaskWithMaintenanceWindowCmd.Flags().String("targets", "", "The targets (either managed nodes or maintenance window targets).")
	ssm_registerTaskWithMaintenanceWindowCmd.Flags().String("task-arn", "", "The ARN of the task to run.")
	ssm_registerTaskWithMaintenanceWindowCmd.Flags().String("task-invocation-parameters", "", "The parameters that the task should use during execution.")
	ssm_registerTaskWithMaintenanceWindowCmd.Flags().String("task-parameters", "", "The parameters that should be passed to the task when it is run.")
	ssm_registerTaskWithMaintenanceWindowCmd.Flags().String("task-type", "", "The type of task being registered.")
	ssm_registerTaskWithMaintenanceWindowCmd.Flags().String("window-id", "", "The ID of the maintenance window the task should be added to.")
	ssm_registerTaskWithMaintenanceWindowCmd.MarkFlagRequired("task-arn")
	ssm_registerTaskWithMaintenanceWindowCmd.MarkFlagRequired("task-type")
	ssm_registerTaskWithMaintenanceWindowCmd.MarkFlagRequired("window-id")
	ssmCmd.AddCommand(ssm_registerTaskWithMaintenanceWindowCmd)
}
