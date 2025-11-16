package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var globalaccelerator_deleteCustomRoutingListenerCmd = &cobra.Command{
	Use:   "delete-custom-routing-listener",
	Short: "Delete a listener for a custom routing accelerator.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(globalaccelerator_deleteCustomRoutingListenerCmd).Standalone()

	globalaccelerator_deleteCustomRoutingListenerCmd.Flags().String("listener-arn", "", "The Amazon Resource Name (ARN) of the listener to delete.")
	globalaccelerator_deleteCustomRoutingListenerCmd.MarkFlagRequired("listener-arn")
	globalacceleratorCmd.AddCommand(globalaccelerator_deleteCustomRoutingListenerCmd)
}
