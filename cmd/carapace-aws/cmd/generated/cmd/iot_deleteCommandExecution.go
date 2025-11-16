package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_deleteCommandExecutionCmd = &cobra.Command{
	Use:   "delete-command-execution",
	Short: "Delete a command execution.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_deleteCommandExecutionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iot_deleteCommandExecutionCmd).Standalone()

		iot_deleteCommandExecutionCmd.Flags().String("execution-id", "", "The unique identifier of the command execution that you want to delete from your account.")
		iot_deleteCommandExecutionCmd.Flags().String("target-arn", "", "The Amazon Resource Number (ARN) of the target device for which you want to delete command executions.")
		iot_deleteCommandExecutionCmd.MarkFlagRequired("execution-id")
		iot_deleteCommandExecutionCmd.MarkFlagRequired("target-arn")
	})
	iotCmd.AddCommand(iot_deleteCommandExecutionCmd)
}
