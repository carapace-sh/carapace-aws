package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var autoscaling_describeScheduledActionsCmd = &cobra.Command{
	Use:   "describe-scheduled-actions",
	Short: "Gets information about the scheduled actions that haven't run or that have not reached their end time.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(autoscaling_describeScheduledActionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(autoscaling_describeScheduledActionsCmd).Standalone()

		autoscaling_describeScheduledActionsCmd.Flags().String("auto-scaling-group-name", "", "The name of the Auto Scaling group.")
		autoscaling_describeScheduledActionsCmd.Flags().String("end-time", "", "The latest scheduled start time to return.")
		autoscaling_describeScheduledActionsCmd.Flags().String("max-records", "", "The maximum number of items to return with this call.")
		autoscaling_describeScheduledActionsCmd.Flags().String("next-token", "", "The token for the next set of items to return.")
		autoscaling_describeScheduledActionsCmd.Flags().String("scheduled-action-names", "", "The names of one or more scheduled actions.")
		autoscaling_describeScheduledActionsCmd.Flags().String("start-time", "", "The earliest scheduled start time to return.")
	})
	autoscalingCmd.AddCommand(autoscaling_describeScheduledActionsCmd)
}
