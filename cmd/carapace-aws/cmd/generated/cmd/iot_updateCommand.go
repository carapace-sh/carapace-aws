package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_updateCommandCmd = &cobra.Command{
	Use:   "update-command",
	Short: "Update information about a command or mark a command for deprecation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_updateCommandCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iot_updateCommandCmd).Standalone()

		iot_updateCommandCmd.Flags().String("command-id", "", "The unique identifier of the command to be updated.")
		iot_updateCommandCmd.Flags().String("deprecated", "", "A boolean that you can use to specify whether to deprecate a command.")
		iot_updateCommandCmd.Flags().String("description", "", "A short text description of the command.")
		iot_updateCommandCmd.Flags().String("display-name", "", "The new user-friendly name to use in the console for the command.")
		iot_updateCommandCmd.MarkFlagRequired("command-id")
	})
	iotCmd.AddCommand(iot_updateCommandCmd)
}
