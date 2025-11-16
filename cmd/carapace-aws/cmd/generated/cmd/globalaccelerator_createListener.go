package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var globalaccelerator_createListenerCmd = &cobra.Command{
	Use:   "create-listener",
	Short: "Create a listener to process inbound connections from clients to an accelerator.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(globalaccelerator_createListenerCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(globalaccelerator_createListenerCmd).Standalone()

		globalaccelerator_createListenerCmd.Flags().String("accelerator-arn", "", "The Amazon Resource Name (ARN) of your accelerator.")
		globalaccelerator_createListenerCmd.Flags().String("client-affinity", "", "Client affinity lets you direct all requests from a user to the same endpoint, if you have stateful applications, regardless of the port and protocol of the client request.")
		globalaccelerator_createListenerCmd.Flags().String("idempotency-token", "", "A unique, case-sensitive identifier that you provide to ensure the idempotency—that is, the uniqueness—of the request.")
		globalaccelerator_createListenerCmd.Flags().String("port-ranges", "", "The list of port ranges to support for connections from clients to your accelerator.")
		globalaccelerator_createListenerCmd.Flags().String("protocol", "", "The protocol for connections from clients to your accelerator.")
		globalaccelerator_createListenerCmd.MarkFlagRequired("accelerator-arn")
		globalaccelerator_createListenerCmd.MarkFlagRequired("idempotency-token")
		globalaccelerator_createListenerCmd.MarkFlagRequired("port-ranges")
		globalaccelerator_createListenerCmd.MarkFlagRequired("protocol")
	})
	globalacceleratorCmd.AddCommand(globalaccelerator_createListenerCmd)
}
