package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var globalaccelerator_removeCustomRoutingEndpointsCmd = &cobra.Command{
	Use:   "remove-custom-routing-endpoints",
	Short: "Remove endpoints from a custom routing accelerator.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(globalaccelerator_removeCustomRoutingEndpointsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(globalaccelerator_removeCustomRoutingEndpointsCmd).Standalone()

		globalaccelerator_removeCustomRoutingEndpointsCmd.Flags().String("endpoint-group-arn", "", "The Amazon Resource Name (ARN) of the endpoint group to remove endpoints from.")
		globalaccelerator_removeCustomRoutingEndpointsCmd.Flags().String("endpoint-ids", "", "The IDs for the endpoints.")
		globalaccelerator_removeCustomRoutingEndpointsCmd.MarkFlagRequired("endpoint-group-arn")
		globalaccelerator_removeCustomRoutingEndpointsCmd.MarkFlagRequired("endpoint-ids")
	})
	globalacceleratorCmd.AddCommand(globalaccelerator_removeCustomRoutingEndpointsCmd)
}
