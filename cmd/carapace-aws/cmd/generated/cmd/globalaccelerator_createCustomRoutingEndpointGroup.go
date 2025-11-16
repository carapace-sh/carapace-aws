package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var globalaccelerator_createCustomRoutingEndpointGroupCmd = &cobra.Command{
	Use:   "create-custom-routing-endpoint-group",
	Short: "Create an endpoint group for the specified listener for a custom routing accelerator.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(globalaccelerator_createCustomRoutingEndpointGroupCmd).Standalone()

	globalaccelerator_createCustomRoutingEndpointGroupCmd.Flags().String("destination-configurations", "", "Sets the port range and protocol for all endpoints (virtual private cloud subnets) in a custom routing endpoint group to accept client traffic on.")
	globalaccelerator_createCustomRoutingEndpointGroupCmd.Flags().String("endpoint-group-region", "", "The Amazon Web Services Region where the endpoint group is located.")
	globalaccelerator_createCustomRoutingEndpointGroupCmd.Flags().String("idempotency-token", "", "A unique, case-sensitive identifier that you provide to ensure the idempotency—that is, the uniqueness—of the request.")
	globalaccelerator_createCustomRoutingEndpointGroupCmd.Flags().String("listener-arn", "", "The Amazon Resource Name (ARN) of the listener for a custom routing endpoint.")
	globalaccelerator_createCustomRoutingEndpointGroupCmd.MarkFlagRequired("destination-configurations")
	globalaccelerator_createCustomRoutingEndpointGroupCmd.MarkFlagRequired("endpoint-group-region")
	globalaccelerator_createCustomRoutingEndpointGroupCmd.MarkFlagRequired("idempotency-token")
	globalaccelerator_createCustomRoutingEndpointGroupCmd.MarkFlagRequired("listener-arn")
	globalacceleratorCmd.AddCommand(globalaccelerator_createCustomRoutingEndpointGroupCmd)
}
