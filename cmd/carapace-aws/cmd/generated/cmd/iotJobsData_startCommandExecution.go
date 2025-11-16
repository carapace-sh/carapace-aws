package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotJobsData_startCommandExecutionCmd = &cobra.Command{
	Use:   "start-command-execution",
	Short: "Using the command created with the `CreateCommand` API, start a command execution on a specific device.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotJobsData_startCommandExecutionCmd).Standalone()

	iotJobsData_startCommandExecutionCmd.Flags().String("client-token", "", "The client token is used to implement idempotency.")
	iotJobsData_startCommandExecutionCmd.Flags().String("command-arn", "", "The Amazon Resource Number (ARN) of the command.")
	iotJobsData_startCommandExecutionCmd.Flags().String("execution-timeout-seconds", "", "Specifies the amount of time in second the device has to finish the command execution.")
	iotJobsData_startCommandExecutionCmd.Flags().String("parameters", "", "A list of parameters that are required by the `StartCommandExecution` API when performing the command on a device.")
	iotJobsData_startCommandExecutionCmd.Flags().String("target-arn", "", "The Amazon Resource Number (ARN) of the device where the command execution is occurring.")
	iotJobsData_startCommandExecutionCmd.MarkFlagRequired("command-arn")
	iotJobsData_startCommandExecutionCmd.MarkFlagRequired("target-arn")
	iotJobsDataCmd.AddCommand(iotJobsData_startCommandExecutionCmd)
}
