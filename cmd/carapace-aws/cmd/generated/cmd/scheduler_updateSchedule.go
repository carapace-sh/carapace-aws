package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var scheduler_updateScheduleCmd = &cobra.Command{
	Use:   "update-schedule",
	Short: "Updates the specified schedule.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(scheduler_updateScheduleCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(scheduler_updateScheduleCmd).Standalone()

		scheduler_updateScheduleCmd.Flags().String("action-after-completion", "", "Specifies the action that EventBridge Scheduler applies to the schedule after the schedule completes invoking the target.")
		scheduler_updateScheduleCmd.Flags().String("client-token", "", "Unique, case-sensitive identifier you provide to ensure the idempotency of the request.")
		scheduler_updateScheduleCmd.Flags().String("description", "", "The description you specify for the schedule.")
		scheduler_updateScheduleCmd.Flags().String("end-date", "", "The date, in UTC, before which the schedule can invoke its target.")
		scheduler_updateScheduleCmd.Flags().String("flexible-time-window", "", "Allows you to configure a time window during which EventBridge Scheduler invokes the schedule.")
		scheduler_updateScheduleCmd.Flags().String("group-name", "", "The name of the schedule group with which the schedule is associated.")
		scheduler_updateScheduleCmd.Flags().String("kms-key-arn", "", "The ARN for the customer managed KMS key that that you want EventBridge Scheduler to use to encrypt and decrypt your data.")
		scheduler_updateScheduleCmd.Flags().String("name", "", "The name of the schedule that you are updating.")
		scheduler_updateScheduleCmd.Flags().String("schedule-expression", "", "The expression that defines when the schedule runs.")
		scheduler_updateScheduleCmd.Flags().String("schedule-expression-timezone", "", "The timezone in which the scheduling expression is evaluated.")
		scheduler_updateScheduleCmd.Flags().String("start-date", "", "The date, in UTC, after which the schedule can begin invoking its target.")
		scheduler_updateScheduleCmd.Flags().String("state", "", "Specifies whether the schedule is enabled or disabled.")
		scheduler_updateScheduleCmd.Flags().String("target", "", "The schedule target.")
		scheduler_updateScheduleCmd.MarkFlagRequired("flexible-time-window")
		scheduler_updateScheduleCmd.MarkFlagRequired("name")
		scheduler_updateScheduleCmd.MarkFlagRequired("schedule-expression")
		scheduler_updateScheduleCmd.MarkFlagRequired("target")
	})
	schedulerCmd.AddCommand(scheduler_updateScheduleCmd)
}
