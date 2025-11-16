package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var stepfunctions_describeStateMachineCmd = &cobra.Command{
	Use:   "describe-state-machine",
	Short: "Provides information about a state machine's definition, its IAM role Amazon Resource Name (ARN), and configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(stepfunctions_describeStateMachineCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(stepfunctions_describeStateMachineCmd).Standalone()

		stepfunctions_describeStateMachineCmd.Flags().String("included-data", "", "If your state machine definition is encrypted with a KMS key, callers must have `kms:Decrypt` permission to decrypt the definition.")
		stepfunctions_describeStateMachineCmd.Flags().String("state-machine-arn", "", "The Amazon Resource Name (ARN) of the state machine for which you want the information.")
		stepfunctions_describeStateMachineCmd.MarkFlagRequired("state-machine-arn")
	})
	stepfunctionsCmd.AddCommand(stepfunctions_describeStateMachineCmd)
}
