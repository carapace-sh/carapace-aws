package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var globalaccelerator_createCustomRoutingListenerCmd = &cobra.Command{
	Use:   "create-custom-routing-listener",
	Short: "Create a listener to process inbound connections from clients to a custom routing accelerator.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(globalaccelerator_createCustomRoutingListenerCmd).Standalone()

	globalaccelerator_createCustomRoutingListenerCmd.Flags().String("accelerator-arn", "", "The Amazon Resource Name (ARN) of the accelerator for a custom routing listener.")
	globalaccelerator_createCustomRoutingListenerCmd.Flags().String("idempotency-token", "", "A unique, case-sensitive identifier that you provide to ensure the idempotency—that is, the uniqueness—of the request.")
	globalaccelerator_createCustomRoutingListenerCmd.Flags().String("port-ranges", "", "The port range to support for connections from clients to your accelerator.")
	globalaccelerator_createCustomRoutingListenerCmd.MarkFlagRequired("accelerator-arn")
	globalaccelerator_createCustomRoutingListenerCmd.MarkFlagRequired("idempotency-token")
	globalaccelerator_createCustomRoutingListenerCmd.MarkFlagRequired("port-ranges")
	globalacceleratorCmd.AddCommand(globalaccelerator_createCustomRoutingListenerCmd)
}
