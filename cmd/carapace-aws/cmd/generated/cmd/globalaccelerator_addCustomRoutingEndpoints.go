package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var globalaccelerator_addCustomRoutingEndpointsCmd = &cobra.Command{
	Use:   "add-custom-routing-endpoints",
	Short: "Associate a virtual private cloud (VPC) subnet endpoint with your custom routing accelerator.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(globalaccelerator_addCustomRoutingEndpointsCmd).Standalone()

	globalaccelerator_addCustomRoutingEndpointsCmd.Flags().String("endpoint-configurations", "", "The list of endpoint objects to add to a custom routing accelerator.")
	globalaccelerator_addCustomRoutingEndpointsCmd.Flags().String("endpoint-group-arn", "", "The Amazon Resource Name (ARN) of the endpoint group for the custom routing endpoint.")
	globalaccelerator_addCustomRoutingEndpointsCmd.MarkFlagRequired("endpoint-configurations")
	globalaccelerator_addCustomRoutingEndpointsCmd.MarkFlagRequired("endpoint-group-arn")
	globalacceleratorCmd.AddCommand(globalaccelerator_addCustomRoutingEndpointsCmd)
}
