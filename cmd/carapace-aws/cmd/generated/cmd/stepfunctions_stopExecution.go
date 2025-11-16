package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var stepfunctions_stopExecutionCmd = &cobra.Command{
	Use:   "stop-execution",
	Short: "Stops an execution.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(stepfunctions_stopExecutionCmd).Standalone()

	stepfunctions_stopExecutionCmd.Flags().String("cause", "", "A more detailed explanation of the cause of the failure.")
	stepfunctions_stopExecutionCmd.Flags().String("error", "", "The error code of the failure.")
	stepfunctions_stopExecutionCmd.Flags().String("execution-arn", "", "The Amazon Resource Name (ARN) of the execution to stop.")
	stepfunctions_stopExecutionCmd.MarkFlagRequired("execution-arn")
	stepfunctionsCmd.AddCommand(stepfunctions_stopExecutionCmd)
}
