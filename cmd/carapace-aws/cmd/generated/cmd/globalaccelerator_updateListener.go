package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var globalaccelerator_updateListenerCmd = &cobra.Command{
	Use:   "update-listener",
	Short: "Update a listener.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(globalaccelerator_updateListenerCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(globalaccelerator_updateListenerCmd).Standalone()

		globalaccelerator_updateListenerCmd.Flags().String("client-affinity", "", "Client affinity lets you direct all requests from a user to the same endpoint, if you have stateful applications, regardless of the port and protocol of the client request.")
		globalaccelerator_updateListenerCmd.Flags().String("listener-arn", "", "The Amazon Resource Name (ARN) of the listener to update.")
		globalaccelerator_updateListenerCmd.Flags().String("port-ranges", "", "The updated list of port ranges for the connections from clients to the accelerator.")
		globalaccelerator_updateListenerCmd.Flags().String("protocol", "", "The updated protocol for the connections from clients to the accelerator.")
		globalaccelerator_updateListenerCmd.MarkFlagRequired("listener-arn")
	})
	globalacceleratorCmd.AddCommand(globalaccelerator_updateListenerCmd)
}
