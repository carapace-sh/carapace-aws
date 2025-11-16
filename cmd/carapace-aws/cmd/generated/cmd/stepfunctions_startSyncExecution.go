package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var stepfunctions_startSyncExecutionCmd = &cobra.Command{
	Use:   "start-sync-execution",
	Short: "Starts a Synchronous Express state machine execution.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(stepfunctions_startSyncExecutionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(stepfunctions_startSyncExecutionCmd).Standalone()

		stepfunctions_startSyncExecutionCmd.Flags().String("included-data", "", "If your state machine definition is encrypted with a KMS key, callers must have `kms:Decrypt` permission to decrypt the definition.")
		stepfunctions_startSyncExecutionCmd.Flags().String("input", "", "The string that contains the JSON input data for the execution, for example:")
		stepfunctions_startSyncExecutionCmd.Flags().String("name", "", "The name of the execution.")
		stepfunctions_startSyncExecutionCmd.Flags().String("state-machine-arn", "", "The Amazon Resource Name (ARN) of the state machine to execute.")
		stepfunctions_startSyncExecutionCmd.Flags().String("trace-header", "", "Passes the X-Ray trace header.")
		stepfunctions_startSyncExecutionCmd.MarkFlagRequired("state-machine-arn")
	})
	stepfunctionsCmd.AddCommand(stepfunctions_startSyncExecutionCmd)
}
