package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var globalaccelerator_listCustomRoutingPortMappingsCmd = &cobra.Command{
	Use:   "list-custom-routing-port-mappings",
	Short: "Provides a complete mapping from the public accelerator IP address and port to destination EC2 instance IP addresses and ports in the virtual public cloud (VPC) subnet endpoint for a custom routing accelerator.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(globalaccelerator_listCustomRoutingPortMappingsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(globalaccelerator_listCustomRoutingPortMappingsCmd).Standalone()

		globalaccelerator_listCustomRoutingPortMappingsCmd.Flags().String("accelerator-arn", "", "The Amazon Resource Name (ARN) of the accelerator to list the custom routing port mappings for.")
		globalaccelerator_listCustomRoutingPortMappingsCmd.Flags().String("endpoint-group-arn", "", "The Amazon Resource Name (ARN) of the endpoint group to list the custom routing port mappings for.")
		globalaccelerator_listCustomRoutingPortMappingsCmd.Flags().String("max-results", "", "The number of destination port mappings that you want to return with this call.")
		globalaccelerator_listCustomRoutingPortMappingsCmd.Flags().String("next-token", "", "The token for the next set of results.")
		globalaccelerator_listCustomRoutingPortMappingsCmd.MarkFlagRequired("accelerator-arn")
	})
	globalacceleratorCmd.AddCommand(globalaccelerator_listCustomRoutingPortMappingsCmd)
}
