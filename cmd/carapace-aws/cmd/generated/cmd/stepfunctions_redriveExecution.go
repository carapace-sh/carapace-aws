package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var stepfunctions_redriveExecutionCmd = &cobra.Command{
	Use:   "redrive-execution",
	Short: "Restarts unsuccessful executions of Standard workflows that didn't complete successfully in the last 14 days.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(stepfunctions_redriveExecutionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(stepfunctions_redriveExecutionCmd).Standalone()

		stepfunctions_redriveExecutionCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
		stepfunctions_redriveExecutionCmd.Flags().String("execution-arn", "", "The Amazon Resource Name (ARN) of the execution to be redriven.")
		stepfunctions_redriveExecutionCmd.MarkFlagRequired("execution-arn")
	})
	stepfunctionsCmd.AddCommand(stepfunctions_redriveExecutionCmd)
}
