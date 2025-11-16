package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssm_updateMaintenanceWindowTaskCmd = &cobra.Command{
	Use:   "update-maintenance-window-task",
	Short: "Modifies a task assigned to a maintenance window.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssm_updateMaintenanceWindowTaskCmd).Standalone()

	ssm_updateMaintenanceWindowTaskCmd.Flags().String("alarm-configuration", "", "The CloudWatch alarm you want to apply to your maintenance window task.")
	ssm_updateMaintenanceWindowTaskCmd.Flags().String("cutoff-behavior", "", "Indicates whether tasks should continue to run after the cutoff time specified in the maintenance windows is reached.")
	ssm_updateMaintenanceWindowTaskCmd.Flags().String("description", "", "The new task description to specify.")
	ssm_updateMaintenanceWindowTaskCmd.Flags().String("logging-info", "", "The new logging location in Amazon S3 to specify.")
	ssm_updateMaintenanceWindowTaskCmd.Flags().String("max-concurrency", "", "The new `MaxConcurrency` value you want to specify.")
	ssm_updateMaintenanceWindowTaskCmd.Flags().String("max-errors", "", "The new `MaxErrors` value to specify.")
	ssm_updateMaintenanceWindowTaskCmd.Flags().String("name", "", "The new task name to specify.")
	ssm_updateMaintenanceWindowTaskCmd.Flags().Bool("no-replace", false, "If True, then all fields that are required by the [RegisterTaskWithMaintenanceWindow]() operation are also required for this API request.")
	ssm_updateMaintenanceWindowTaskCmd.Flags().String("priority", "", "The new task priority to specify.")
	ssm_updateMaintenanceWindowTaskCmd.Flags().Bool("replace", false, "If True, then all fields that are required by the [RegisterTaskWithMaintenanceWindow]() operation are also required for this API request.")
	ssm_updateMaintenanceWindowTaskCmd.Flags().String("service-role-arn", "", "The Amazon Resource Name (ARN) of the IAM service role for Amazon Web Services Systems Manager to assume when running a maintenance window task.")
	ssm_updateMaintenanceWindowTaskCmd.Flags().String("targets", "", "The targets (either managed nodes or tags) to modify.")
	ssm_updateMaintenanceWindowTaskCmd.Flags().String("task-arn", "", "The task ARN to modify.")
	ssm_updateMaintenanceWindowTaskCmd.Flags().String("task-invocation-parameters", "", "The parameters that the task should use during execution.")
	ssm_updateMaintenanceWindowTaskCmd.Flags().String("task-parameters", "", "The parameters to modify.")
	ssm_updateMaintenanceWindowTaskCmd.Flags().String("window-id", "", "The maintenance window ID that contains the task to modify.")
	ssm_updateMaintenanceWindowTaskCmd.Flags().String("window-task-id", "", "The task ID to modify.")
	ssm_updateMaintenanceWindowTaskCmd.Flag("no-replace").Hidden = true
	ssm_updateMaintenanceWindowTaskCmd.MarkFlagRequired("window-id")
	ssm_updateMaintenanceWindowTaskCmd.MarkFlagRequired("window-task-id")
	ssmCmd.AddCommand(ssm_updateMaintenanceWindowTaskCmd)
}
