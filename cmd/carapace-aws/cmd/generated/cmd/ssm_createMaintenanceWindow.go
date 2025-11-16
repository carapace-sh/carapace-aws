package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssm_createMaintenanceWindowCmd = &cobra.Command{
	Use:   "create-maintenance-window",
	Short: "Creates a new maintenance window.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssm_createMaintenanceWindowCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ssm_createMaintenanceWindowCmd).Standalone()

		ssm_createMaintenanceWindowCmd.Flags().String("allow-unassociated-targets", "", "Enables a maintenance window task to run on managed nodes, even if you haven't registered those nodes as targets.")
		ssm_createMaintenanceWindowCmd.Flags().String("client-token", "", "User-provided idempotency token.")
		ssm_createMaintenanceWindowCmd.Flags().String("cutoff", "", "The number of hours before the end of the maintenance window that Amazon Web Services Systems Manager stops scheduling new tasks for execution.")
		ssm_createMaintenanceWindowCmd.Flags().String("description", "", "An optional description for the maintenance window.")
		ssm_createMaintenanceWindowCmd.Flags().String("duration", "", "The duration of the maintenance window in hours.")
		ssm_createMaintenanceWindowCmd.Flags().String("end-date", "", "The date and time, in ISO-8601 Extended format, for when you want the maintenance window to become inactive.")
		ssm_createMaintenanceWindowCmd.Flags().String("name", "", "The name of the maintenance window.")
		ssm_createMaintenanceWindowCmd.Flags().String("schedule", "", "The schedule of the maintenance window in the form of a cron or rate expression.")
		ssm_createMaintenanceWindowCmd.Flags().String("schedule-offset", "", "The number of days to wait after the date and time specified by a cron expression before running the maintenance window.")
		ssm_createMaintenanceWindowCmd.Flags().String("schedule-timezone", "", "The time zone that the scheduled maintenance window executions are based on, in Internet Assigned Numbers Authority (IANA) format.")
		ssm_createMaintenanceWindowCmd.Flags().String("start-date", "", "The date and time, in ISO-8601 Extended format, for when you want the maintenance window to become active.")
		ssm_createMaintenanceWindowCmd.Flags().String("tags", "", "Optional metadata that you assign to a resource.")
		ssm_createMaintenanceWindowCmd.MarkFlagRequired("allow-unassociated-targets")
		ssm_createMaintenanceWindowCmd.MarkFlagRequired("cutoff")
		ssm_createMaintenanceWindowCmd.MarkFlagRequired("duration")
		ssm_createMaintenanceWindowCmd.MarkFlagRequired("name")
		ssm_createMaintenanceWindowCmd.MarkFlagRequired("schedule")
	})
	ssmCmd.AddCommand(ssm_createMaintenanceWindowCmd)
}
