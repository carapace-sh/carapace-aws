package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var globalaccelerator_deleteCustomRoutingEndpointGroupCmd = &cobra.Command{
	Use:   "delete-custom-routing-endpoint-group",
	Short: "Delete an endpoint group from a listener for a custom routing accelerator.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(globalaccelerator_deleteCustomRoutingEndpointGroupCmd).Standalone()

	globalaccelerator_deleteCustomRoutingEndpointGroupCmd.Flags().String("endpoint-group-arn", "", "The Amazon Resource Name (ARN) of the endpoint group to delete.")
	globalaccelerator_deleteCustomRoutingEndpointGroupCmd.MarkFlagRequired("endpoint-group-arn")
	globalacceleratorCmd.AddCommand(globalaccelerator_deleteCustomRoutingEndpointGroupCmd)
}
