package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssm_updateMaintenanceWindowCmd = &cobra.Command{
	Use:   "update-maintenance-window",
	Short: "Updates an existing maintenance window.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssm_updateMaintenanceWindowCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ssm_updateMaintenanceWindowCmd).Standalone()

		ssm_updateMaintenanceWindowCmd.Flags().String("allow-unassociated-targets", "", "Whether targets must be registered with the maintenance window before tasks can be defined for those targets.")
		ssm_updateMaintenanceWindowCmd.Flags().String("cutoff", "", "The number of hours before the end of the maintenance window that Amazon Web Services Systems Manager stops scheduling new tasks for execution.")
		ssm_updateMaintenanceWindowCmd.Flags().String("description", "", "An optional description for the update request.")
		ssm_updateMaintenanceWindowCmd.Flags().String("duration", "", "The duration of the maintenance window in hours.")
		ssm_updateMaintenanceWindowCmd.Flags().String("enabled", "", "Whether the maintenance window is enabled.")
		ssm_updateMaintenanceWindowCmd.Flags().String("end-date", "", "The date and time, in ISO-8601 Extended format, for when you want the maintenance window to become inactive.")
		ssm_updateMaintenanceWindowCmd.Flags().String("name", "", "The name of the maintenance window.")
		ssm_updateMaintenanceWindowCmd.Flags().Bool("no-replace", false, "If `True`, then all fields that are required by the [CreateMaintenanceWindow]() operation are also required for this API request.")
		ssm_updateMaintenanceWindowCmd.Flags().Bool("replace", false, "If `True`, then all fields that are required by the [CreateMaintenanceWindow]() operation are also required for this API request.")
		ssm_updateMaintenanceWindowCmd.Flags().String("schedule", "", "The schedule of the maintenance window in the form of a cron or rate expression.")
		ssm_updateMaintenanceWindowCmd.Flags().String("schedule-offset", "", "The number of days to wait after the date and time specified by a cron expression before running the maintenance window.")
		ssm_updateMaintenanceWindowCmd.Flags().String("schedule-timezone", "", "The time zone that the scheduled maintenance window executions are based on, in Internet Assigned Numbers Authority (IANA) format.")
		ssm_updateMaintenanceWindowCmd.Flags().String("start-date", "", "The date and time, in ISO-8601 Extended format, for when you want the maintenance window to become active.")
		ssm_updateMaintenanceWindowCmd.Flags().String("window-id", "", "The ID of the maintenance window to update.")
		ssm_updateMaintenanceWindowCmd.Flag("no-replace").Hidden = true
		ssm_updateMaintenanceWindowCmd.MarkFlagRequired("window-id")
	})
	ssmCmd.AddCommand(ssm_updateMaintenanceWindowCmd)
}
