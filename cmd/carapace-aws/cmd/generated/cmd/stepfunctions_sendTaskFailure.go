package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var stepfunctions_sendTaskFailureCmd = &cobra.Command{
	Use:   "send-task-failure",
	Short: "Used by activity workers, Task states using the [callback](https://docs.aws.amazon.com/step-functions/latest/dg/connect-to-resource.html#connect-wait-token) pattern, and optionally Task states using the [job run](https://docs.aws.amazon.com/step-functions/latest/dg/connect-to-resource.html#connect-sync) pattern to report that the task identified by the `taskToken` failed.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(stepfunctions_sendTaskFailureCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(stepfunctions_sendTaskFailureCmd).Standalone()

		stepfunctions_sendTaskFailureCmd.Flags().String("cause", "", "A more detailed explanation of the cause of the failure.")
		stepfunctions_sendTaskFailureCmd.Flags().String("error", "", "The error code of the failure.")
		stepfunctions_sendTaskFailureCmd.Flags().String("task-token", "", "The token that represents this task.")
		stepfunctions_sendTaskFailureCmd.MarkFlagRequired("task-token")
	})
	stepfunctionsCmd.AddCommand(stepfunctions_sendTaskFailureCmd)
}
