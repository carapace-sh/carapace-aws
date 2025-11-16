package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var globalaccelerator_listCustomRoutingPortMappingsByDestinationCmd = &cobra.Command{
	Use:   "list-custom-routing-port-mappings-by-destination",
	Short: "List the port mappings for a specific EC2 instance (destination) in a VPC subnet endpoint.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(globalaccelerator_listCustomRoutingPortMappingsByDestinationCmd).Standalone()

	globalaccelerator_listCustomRoutingPortMappingsByDestinationCmd.Flags().String("destination-address", "", "The endpoint IP address in a virtual private cloud (VPC) subnet for which you want to receive back port mappings.")
	globalaccelerator_listCustomRoutingPortMappingsByDestinationCmd.Flags().String("endpoint-id", "", "The ID for the virtual private cloud (VPC) subnet.")
	globalaccelerator_listCustomRoutingPortMappingsByDestinationCmd.Flags().String("max-results", "", "The number of destination port mappings that you want to return with this call.")
	globalaccelerator_listCustomRoutingPortMappingsByDestinationCmd.Flags().String("next-token", "", "The token for the next set of results.")
	globalaccelerator_listCustomRoutingPortMappingsByDestinationCmd.MarkFlagRequired("destination-address")
	globalaccelerator_listCustomRoutingPortMappingsByDestinationCmd.MarkFlagRequired("endpoint-id")
	globalacceleratorCmd.AddCommand(globalaccelerator_listCustomRoutingPortMappingsByDestinationCmd)
}
