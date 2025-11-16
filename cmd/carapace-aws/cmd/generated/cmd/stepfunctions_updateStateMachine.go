package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var stepfunctions_updateStateMachineCmd = &cobra.Command{
	Use:   "update-state-machine",
	Short: "Updates an existing state machine by modifying its `definition`, `roleArn`, `loggingConfiguration`, or `EncryptionConfiguration`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(stepfunctions_updateStateMachineCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(stepfunctions_updateStateMachineCmd).Standalone()

		stepfunctions_updateStateMachineCmd.Flags().String("definition", "", "The Amazon States Language definition of the state machine.")
		stepfunctions_updateStateMachineCmd.Flags().String("encryption-configuration", "", "Settings to configure server-side encryption.")
		stepfunctions_updateStateMachineCmd.Flags().String("logging-configuration", "", "Use the `LoggingConfiguration` data type to set CloudWatch Logs options.")
		stepfunctions_updateStateMachineCmd.Flags().String("publish", "", "Specifies whether the state machine version is published.")
		stepfunctions_updateStateMachineCmd.Flags().String("role-arn", "", "The Amazon Resource Name (ARN) of the IAM role of the state machine.")
		stepfunctions_updateStateMachineCmd.Flags().String("state-machine-arn", "", "The Amazon Resource Name (ARN) of the state machine.")
		stepfunctions_updateStateMachineCmd.Flags().String("tracing-configuration", "", "Selects whether X-Ray tracing is enabled.")
		stepfunctions_updateStateMachineCmd.Flags().String("version-description", "", "An optional description of the state machine version to publish.")
		stepfunctions_updateStateMachineCmd.MarkFlagRequired("state-machine-arn")
	})
	stepfunctionsCmd.AddCommand(stepfunctions_updateStateMachineCmd)
}
