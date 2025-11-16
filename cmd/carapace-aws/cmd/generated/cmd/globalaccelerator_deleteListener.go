package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var globalaccelerator_deleteListenerCmd = &cobra.Command{
	Use:   "delete-listener",
	Short: "Delete a listener from an accelerator.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(globalaccelerator_deleteListenerCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(globalaccelerator_deleteListenerCmd).Standalone()

		globalaccelerator_deleteListenerCmd.Flags().String("listener-arn", "", "The Amazon Resource Name (ARN) of the listener.")
		globalaccelerator_deleteListenerCmd.MarkFlagRequired("listener-arn")
	})
	globalacceleratorCmd.AddCommand(globalaccelerator_deleteListenerCmd)
}
