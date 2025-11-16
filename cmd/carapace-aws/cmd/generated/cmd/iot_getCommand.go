package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_getCommandCmd = &cobra.Command{
	Use:   "get-command",
	Short: "Gets information about the specified command.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_getCommandCmd).Standalone()

	iot_getCommandCmd.Flags().String("command-id", "", "The unique identifier of the command for which you want to retrieve information.")
	iot_getCommandCmd.MarkFlagRequired("command-id")
	iotCmd.AddCommand(iot_getCommandCmd)
}
