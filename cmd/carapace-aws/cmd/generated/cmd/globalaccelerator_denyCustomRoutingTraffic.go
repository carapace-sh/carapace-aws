package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var globalaccelerator_denyCustomRoutingTrafficCmd = &cobra.Command{
	Use:   "deny-custom-routing-traffic",
	Short: "Specify the Amazon EC2 instance (destination) IP addresses and ports for a VPC subnet endpoint that cannot receive traffic for a custom routing accelerator.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(globalaccelerator_denyCustomRoutingTrafficCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(globalaccelerator_denyCustomRoutingTrafficCmd).Standalone()

		globalaccelerator_denyCustomRoutingTrafficCmd.Flags().String("deny-all-traffic-to-endpoint", "", "Indicates whether all destination IP addresses and ports for a specified VPC subnet endpoint *cannot* receive traffic from a custom routing accelerator.")
		globalaccelerator_denyCustomRoutingTrafficCmd.Flags().String("destination-addresses", "", "A list of specific Amazon EC2 instance IP addresses (destination addresses) in a subnet that you want to prevent from receiving traffic.")
		globalaccelerator_denyCustomRoutingTrafficCmd.Flags().String("destination-ports", "", "A list of specific Amazon EC2 instance ports (destination ports) in a subnet endpoint that you want to prevent from receiving traffic.")
		globalaccelerator_denyCustomRoutingTrafficCmd.Flags().String("endpoint-group-arn", "", "The Amazon Resource Name (ARN) of the endpoint group.")
		globalaccelerator_denyCustomRoutingTrafficCmd.Flags().String("endpoint-id", "", "An ID for the endpoint.")
		globalaccelerator_denyCustomRoutingTrafficCmd.MarkFlagRequired("endpoint-group-arn")
		globalaccelerator_denyCustomRoutingTrafficCmd.MarkFlagRequired("endpoint-id")
	})
	globalacceleratorCmd.AddCommand(globalaccelerator_denyCustomRoutingTrafficCmd)
}
