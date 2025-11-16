package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var stepfunctions_getActivityTaskCmd = &cobra.Command{
	Use:   "get-activity-task",
	Short: "Used by workers to retrieve a task (with the specified activity ARN) which has been scheduled for execution by a running state machine.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(stepfunctions_getActivityTaskCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(stepfunctions_getActivityTaskCmd).Standalone()

		stepfunctions_getActivityTaskCmd.Flags().String("activity-arn", "", "The Amazon Resource Name (ARN) of the activity to retrieve tasks from (assigned when you create the task using [CreateActivity]().)")
		stepfunctions_getActivityTaskCmd.Flags().String("worker-name", "", "You can provide an arbitrary name in order to identify the worker that the task is assigned to.")
		stepfunctions_getActivityTaskCmd.MarkFlagRequired("activity-arn")
	})
	stepfunctionsCmd.AddCommand(stepfunctions_getActivityTaskCmd)
}
