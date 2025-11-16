package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var globalaccelerator_describeCustomRoutingEndpointGroupCmd = &cobra.Command{
	Use:   "describe-custom-routing-endpoint-group",
	Short: "Describe an endpoint group for a custom routing accelerator.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(globalaccelerator_describeCustomRoutingEndpointGroupCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(globalaccelerator_describeCustomRoutingEndpointGroupCmd).Standalone()

		globalaccelerator_describeCustomRoutingEndpointGroupCmd.Flags().String("endpoint-group-arn", "", "The Amazon Resource Name (ARN) of the endpoint group to describe.")
		globalaccelerator_describeCustomRoutingEndpointGroupCmd.MarkFlagRequired("endpoint-group-arn")
	})
	globalacceleratorCmd.AddCommand(globalaccelerator_describeCustomRoutingEndpointGroupCmd)
}
