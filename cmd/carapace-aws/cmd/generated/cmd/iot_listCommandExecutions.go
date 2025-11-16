package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_listCommandExecutionsCmd = &cobra.Command{
	Use:   "list-command-executions",
	Short: "List all command executions.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_listCommandExecutionsCmd).Standalone()

	iot_listCommandExecutionsCmd.Flags().String("command-arn", "", "The Amazon Resource Number (ARN) of the command.")
	iot_listCommandExecutionsCmd.Flags().String("completed-time-filter", "", "List all command executions that completed any time before or after the date and time that you specify.")
	iot_listCommandExecutionsCmd.Flags().String("max-results", "", "The maximum number of results to return in this operation.")
	iot_listCommandExecutionsCmd.Flags().String("namespace", "", "The namespace of the command.")
	iot_listCommandExecutionsCmd.Flags().String("next-token", "", "To retrieve the next set of results, the `nextToken` value from a previous response; otherwise `null` to receive the first set of results.")
	iot_listCommandExecutionsCmd.Flags().String("sort-order", "", "Specify whether to list the command executions that were created in the ascending or descending order.")
	iot_listCommandExecutionsCmd.Flags().String("started-time-filter", "", "List all command executions that started any time before or after the date and time that you specify.")
	iot_listCommandExecutionsCmd.Flags().String("status", "", "List all command executions for the device that have a particular status.")
	iot_listCommandExecutionsCmd.Flags().String("target-arn", "", "The Amazon Resource Number (ARN) of the target device.")
	iotCmd.AddCommand(iot_listCommandExecutionsCmd)
}
