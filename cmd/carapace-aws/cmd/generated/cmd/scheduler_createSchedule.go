package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var scheduler_createScheduleCmd = &cobra.Command{
	Use:   "create-schedule",
	Short: "Creates the specified schedule.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(scheduler_createScheduleCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(scheduler_createScheduleCmd).Standalone()

		scheduler_createScheduleCmd.Flags().String("action-after-completion", "", "Specifies the action that EventBridge Scheduler applies to the schedule after the schedule completes invoking the target.")
		scheduler_createScheduleCmd.Flags().String("client-token", "", "Unique, case-sensitive identifier you provide to ensure the idempotency of the request.")
		scheduler_createScheduleCmd.Flags().String("description", "", "The description you specify for the schedule.")
		scheduler_createScheduleCmd.Flags().String("end-date", "", "The date, in UTC, before which the schedule can invoke its target.")
		scheduler_createScheduleCmd.Flags().String("flexible-time-window", "", "Allows you to configure a time window during which EventBridge Scheduler invokes the schedule.")
		scheduler_createScheduleCmd.Flags().String("group-name", "", "The name of the schedule group to associate with this schedule.")
		scheduler_createScheduleCmd.Flags().String("kms-key-arn", "", "The Amazon Resource Name (ARN) for the customer managed KMS key that EventBridge Scheduler will use to encrypt and decrypt your data.")
		scheduler_createScheduleCmd.Flags().String("name", "", "The name of the schedule that you are creating.")
		scheduler_createScheduleCmd.Flags().String("schedule-expression", "", "The expression that defines when the schedule runs.")
		scheduler_createScheduleCmd.Flags().String("schedule-expression-timezone", "", "The timezone in which the scheduling expression is evaluated.")
		scheduler_createScheduleCmd.Flags().String("start-date", "", "The date, in UTC, after which the schedule can begin invoking its target.")
		scheduler_createScheduleCmd.Flags().String("state", "", "Specifies whether the schedule is enabled or disabled.")
		scheduler_createScheduleCmd.Flags().String("target", "", "The schedule's target.")
		scheduler_createScheduleCmd.MarkFlagRequired("flexible-time-window")
		scheduler_createScheduleCmd.MarkFlagRequired("name")
		scheduler_createScheduleCmd.MarkFlagRequired("schedule-expression")
		scheduler_createScheduleCmd.MarkFlagRequired("target")
	})
	schedulerCmd.AddCommand(scheduler_createScheduleCmd)
}
