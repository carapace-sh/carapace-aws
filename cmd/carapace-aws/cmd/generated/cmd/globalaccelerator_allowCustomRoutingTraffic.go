package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var globalaccelerator_allowCustomRoutingTrafficCmd = &cobra.Command{
	Use:   "allow-custom-routing-traffic",
	Short: "Specify the Amazon EC2 instance (destination) IP addresses and ports for a VPC subnet endpoint that can receive traffic for a custom routing accelerator.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(globalaccelerator_allowCustomRoutingTrafficCmd).Standalone()

	globalaccelerator_allowCustomRoutingTrafficCmd.Flags().String("allow-all-traffic-to-endpoint", "", "Indicates whether all destination IP addresses and ports for a specified VPC subnet endpoint can receive traffic from a custom routing accelerator.")
	globalaccelerator_allowCustomRoutingTrafficCmd.Flags().String("destination-addresses", "", "A list of specific Amazon EC2 instance IP addresses (destination addresses) in a subnet that you want to allow to receive traffic.")
	globalaccelerator_allowCustomRoutingTrafficCmd.Flags().String("destination-ports", "", "A list of specific Amazon EC2 instance ports (destination ports) that you want to allow to receive traffic.")
	globalaccelerator_allowCustomRoutingTrafficCmd.Flags().String("endpoint-group-arn", "", "The Amazon Resource Name (ARN) of the endpoint group.")
	globalaccelerator_allowCustomRoutingTrafficCmd.Flags().String("endpoint-id", "", "An ID for the endpoint.")
	globalaccelerator_allowCustomRoutingTrafficCmd.MarkFlagRequired("endpoint-group-arn")
	globalaccelerator_allowCustomRoutingTrafficCmd.MarkFlagRequired("endpoint-id")
	globalacceleratorCmd.AddCommand(globalaccelerator_allowCustomRoutingTrafficCmd)
}
