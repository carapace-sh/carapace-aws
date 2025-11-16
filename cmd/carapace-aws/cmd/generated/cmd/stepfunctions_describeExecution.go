package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var stepfunctions_describeExecutionCmd = &cobra.Command{
	Use:   "describe-execution",
	Short: "Provides information about a state machine execution, such as the state machine associated with the execution, the execution input and output, and relevant execution metadata.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(stepfunctions_describeExecutionCmd).Standalone()

	stepfunctions_describeExecutionCmd.Flags().String("execution-arn", "", "The Amazon Resource Name (ARN) of the execution to describe.")
	stepfunctions_describeExecutionCmd.Flags().String("included-data", "", "If your state machine definition is encrypted with a KMS key, callers must have `kms:Decrypt` permission to decrypt the definition.")
	stepfunctions_describeExecutionCmd.MarkFlagRequired("execution-arn")
	stepfunctionsCmd.AddCommand(stepfunctions_describeExecutionCmd)
}
