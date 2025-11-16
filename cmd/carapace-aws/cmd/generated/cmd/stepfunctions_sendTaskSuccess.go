package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var stepfunctions_sendTaskSuccessCmd = &cobra.Command{
	Use:   "send-task-success",
	Short: "Used by activity workers, Task states using the [callback](https://docs.aws.amazon.com/step-functions/latest/dg/connect-to-resource.html#connect-wait-token) pattern, and optionally Task states using the [job run](https://docs.aws.amazon.com/step-functions/latest/dg/connect-to-resource.html#connect-sync) pattern to report that the task identified by the `taskToken` completed successfully.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(stepfunctions_sendTaskSuccessCmd).Standalone()

	stepfunctions_sendTaskSuccessCmd.Flags().String("output", "", "The JSON output of the task.")
	stepfunctions_sendTaskSuccessCmd.Flags().String("task-token", "", "The token that represents this task.")
	stepfunctions_sendTaskSuccessCmd.MarkFlagRequired("output")
	stepfunctions_sendTaskSuccessCmd.MarkFlagRequired("task-token")
	stepfunctionsCmd.AddCommand(stepfunctions_sendTaskSuccessCmd)
}
