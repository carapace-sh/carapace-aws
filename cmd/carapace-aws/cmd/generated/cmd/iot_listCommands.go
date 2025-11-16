package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_listCommandsCmd = &cobra.Command{
	Use:   "list-commands",
	Short: "List all commands in your account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_listCommandsCmd).Standalone()

	iot_listCommandsCmd.Flags().String("command-parameter-name", "", "A filter that can be used to display the list of commands that have a specific command parameter name.")
	iot_listCommandsCmd.Flags().String("max-results", "", "The maximum number of results to return in this operation.")
	iot_listCommandsCmd.Flags().String("namespace", "", "The namespace of the command.")
	iot_listCommandsCmd.Flags().String("next-token", "", "To retrieve the next set of results, the `nextToken` value from a previous response; otherwise `null` to receive the first set of results.")
	iot_listCommandsCmd.Flags().String("sort-order", "", "Specify whether to list the commands that you have created in the ascending or descending order.")
	iotCmd.AddCommand(iot_listCommandsCmd)
}
