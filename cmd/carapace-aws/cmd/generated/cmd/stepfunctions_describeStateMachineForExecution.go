package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var stepfunctions_describeStateMachineForExecutionCmd = &cobra.Command{
	Use:   "describe-state-machine-for-execution",
	Short: "Provides information about a state machine's definition, its execution role ARN, and configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(stepfunctions_describeStateMachineForExecutionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(stepfunctions_describeStateMachineForExecutionCmd).Standalone()

		stepfunctions_describeStateMachineForExecutionCmd.Flags().String("execution-arn", "", "The Amazon Resource Name (ARN) of the execution you want state machine information for.")
		stepfunctions_describeStateMachineForExecutionCmd.Flags().String("included-data", "", "If your state machine definition is encrypted with a KMS key, callers must have `kms:Decrypt` permission to decrypt the definition.")
		stepfunctions_describeStateMachineForExecutionCmd.MarkFlagRequired("execution-arn")
	})
	stepfunctionsCmd.AddCommand(stepfunctions_describeStateMachineForExecutionCmd)
}
