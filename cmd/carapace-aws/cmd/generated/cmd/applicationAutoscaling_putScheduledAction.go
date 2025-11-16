package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var applicationAutoscaling_putScheduledActionCmd = &cobra.Command{
	Use:   "put-scheduled-action",
	Short: "Creates or updates a scheduled action for an Application Auto Scaling scalable target.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(applicationAutoscaling_putScheduledActionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(applicationAutoscaling_putScheduledActionCmd).Standalone()

		applicationAutoscaling_putScheduledActionCmd.Flags().String("end-time", "", "The date and time for the recurring schedule to end, in UTC.")
		applicationAutoscaling_putScheduledActionCmd.Flags().String("resource-id", "", "The identifier of the resource associated with the scheduled action.")
		applicationAutoscaling_putScheduledActionCmd.Flags().String("scalable-dimension", "", "The scalable dimension.")
		applicationAutoscaling_putScheduledActionCmd.Flags().String("scalable-target-action", "", "The new minimum and maximum capacity.")
		applicationAutoscaling_putScheduledActionCmd.Flags().String("schedule", "", "The schedule for this action.")
		applicationAutoscaling_putScheduledActionCmd.Flags().String("scheduled-action-name", "", "The name of the scheduled action.")
		applicationAutoscaling_putScheduledActionCmd.Flags().String("service-namespace", "", "The namespace of the Amazon Web Services service that provides the resource.")
		applicationAutoscaling_putScheduledActionCmd.Flags().String("start-time", "", "The date and time for this scheduled action to start, in UTC.")
		applicationAutoscaling_putScheduledActionCmd.Flags().String("timezone", "", "Specifies the time zone used when setting a scheduled action by using an at or cron expression.")
		applicationAutoscaling_putScheduledActionCmd.MarkFlagRequired("resource-id")
		applicationAutoscaling_putScheduledActionCmd.MarkFlagRequired("scalable-dimension")
		applicationAutoscaling_putScheduledActionCmd.MarkFlagRequired("scheduled-action-name")
		applicationAutoscaling_putScheduledActionCmd.MarkFlagRequired("service-namespace")
	})
	applicationAutoscalingCmd.AddCommand(applicationAutoscaling_putScheduledActionCmd)
}
