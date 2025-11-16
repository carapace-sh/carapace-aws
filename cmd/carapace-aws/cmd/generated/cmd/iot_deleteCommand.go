package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_deleteCommandCmd = &cobra.Command{
	Use:   "delete-command",
	Short: "Delete a command resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_deleteCommandCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iot_deleteCommandCmd).Standalone()

		iot_deleteCommandCmd.Flags().String("command-id", "", "The unique identifier of the command to be deleted.")
		iot_deleteCommandCmd.MarkFlagRequired("command-id")
	})
	iotCmd.AddCommand(iot_deleteCommandCmd)
}
