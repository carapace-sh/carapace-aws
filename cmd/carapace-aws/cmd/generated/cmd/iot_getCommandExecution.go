package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_getCommandExecutionCmd = &cobra.Command{
	Use:   "get-command-execution",
	Short: "Gets information about the specific command execution on a single device.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_getCommandExecutionCmd).Standalone()

	iot_getCommandExecutionCmd.Flags().String("execution-id", "", "The unique identifier for the command execution.")
	iot_getCommandExecutionCmd.Flags().String("include-result", "", "Can be used to specify whether to include the result of the command execution in the `GetCommandExecution` API response.")
	iot_getCommandExecutionCmd.Flags().String("target-arn", "", "The Amazon Resource Number (ARN) of the device on which the command execution is being performed.")
	iot_getCommandExecutionCmd.MarkFlagRequired("execution-id")
	iot_getCommandExecutionCmd.MarkFlagRequired("target-arn")
	iotCmd.AddCommand(iot_getCommandExecutionCmd)
}
