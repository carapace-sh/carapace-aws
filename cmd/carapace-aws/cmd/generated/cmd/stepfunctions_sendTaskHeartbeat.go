package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var stepfunctions_sendTaskHeartbeatCmd = &cobra.Command{
	Use:   "send-task-heartbeat",
	Short: "Used by activity workers and Task states using the [callback](https://docs.aws.amazon.com/step-functions/latest/dg/connect-to-resource.html#connect-wait-token) pattern, and optionally Task states using the [job run](https://docs.aws.amazon.com/step-functions/latest/dg/connect-to-resource.html#connect-sync) pattern to report to Step Functions that the task represented by the specified `taskToken` is still making progress.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(stepfunctions_sendTaskHeartbeatCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(stepfunctions_sendTaskHeartbeatCmd).Standalone()

		stepfunctions_sendTaskHeartbeatCmd.Flags().String("task-token", "", "The token that represents this task.")
		stepfunctions_sendTaskHeartbeatCmd.MarkFlagRequired("task-token")
	})
	stepfunctionsCmd.AddCommand(stepfunctions_sendTaskHeartbeatCmd)
}
