package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var globalaccelerator_updateCustomRoutingListenerCmd = &cobra.Command{
	Use:   "update-custom-routing-listener",
	Short: "Update a listener for a custom routing accelerator.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(globalaccelerator_updateCustomRoutingListenerCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(globalaccelerator_updateCustomRoutingListenerCmd).Standalone()

		globalaccelerator_updateCustomRoutingListenerCmd.Flags().String("listener-arn", "", "The Amazon Resource Name (ARN) of the listener to update.")
		globalaccelerator_updateCustomRoutingListenerCmd.Flags().String("port-ranges", "", "The updated port range to support for connections from clients to your accelerator.")
		globalaccelerator_updateCustomRoutingListenerCmd.MarkFlagRequired("listener-arn")
		globalaccelerator_updateCustomRoutingListenerCmd.MarkFlagRequired("port-ranges")
	})
	globalacceleratorCmd.AddCommand(globalaccelerator_updateCustomRoutingListenerCmd)
}
