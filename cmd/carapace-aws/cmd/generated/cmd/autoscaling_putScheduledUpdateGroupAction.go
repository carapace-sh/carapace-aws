package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var autoscaling_putScheduledUpdateGroupActionCmd = &cobra.Command{
	Use:   "put-scheduled-update-group-action",
	Short: "Creates or updates a scheduled scaling action for an Auto Scaling group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(autoscaling_putScheduledUpdateGroupActionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(autoscaling_putScheduledUpdateGroupActionCmd).Standalone()

		autoscaling_putScheduledUpdateGroupActionCmd.Flags().String("auto-scaling-group-name", "", "The name of the Auto Scaling group.")
		autoscaling_putScheduledUpdateGroupActionCmd.Flags().String("desired-capacity", "", "The desired capacity is the initial capacity of the Auto Scaling group after the scheduled action runs and the capacity it attempts to maintain.")
		autoscaling_putScheduledUpdateGroupActionCmd.Flags().String("end-time", "", "The date and time for the recurring schedule to end, in UTC.")
		autoscaling_putScheduledUpdateGroupActionCmd.Flags().String("max-size", "", "The maximum size of the Auto Scaling group.")
		autoscaling_putScheduledUpdateGroupActionCmd.Flags().String("min-size", "", "The minimum size of the Auto Scaling group.")
		autoscaling_putScheduledUpdateGroupActionCmd.Flags().String("recurrence", "", "The recurring schedule for this action.")
		autoscaling_putScheduledUpdateGroupActionCmd.Flags().String("scheduled-action-name", "", "The name of this scaling action.")
		autoscaling_putScheduledUpdateGroupActionCmd.Flags().String("start-time", "", "The date and time for this action to start, in YYYY-MM-DDThh:mm:ssZ format in UTC/GMT only and in quotes (for example, `\"2021-06-01T00:00:00Z\"`).")
		autoscaling_putScheduledUpdateGroupActionCmd.Flags().String("time", "", "This property is no longer used.")
		autoscaling_putScheduledUpdateGroupActionCmd.Flags().String("time-zone", "", "Specifies the time zone for a cron expression.")
		autoscaling_putScheduledUpdateGroupActionCmd.MarkFlagRequired("auto-scaling-group-name")
		autoscaling_putScheduledUpdateGroupActionCmd.MarkFlagRequired("scheduled-action-name")
	})
	autoscalingCmd.AddCommand(autoscaling_putScheduledUpdateGroupActionCmd)
}
