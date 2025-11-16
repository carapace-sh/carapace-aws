package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var networkFirewall_listVpcEndpointAssociationsCmd = &cobra.Command{
	Use:   "list-vpc-endpoint-associations",
	Short: "Retrieves the metadata for the VPC endpoint associations that you have defined.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(networkFirewall_listVpcEndpointAssociationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(networkFirewall_listVpcEndpointAssociationsCmd).Standalone()

		networkFirewall_listVpcEndpointAssociationsCmd.Flags().String("firewall-arn", "", "The Amazon Resource Name (ARN) of the firewall.")
		networkFirewall_listVpcEndpointAssociationsCmd.Flags().String("max-results", "", "The maximum number of objects that you want Network Firewall to return for this request.")
		networkFirewall_listVpcEndpointAssociationsCmd.Flags().String("next-token", "", "When you request a list of objects with a `MaxResults` setting, if the number of objects that are still available for retrieval exceeds the maximum you requested, Network Firewall returns a `NextToken` value in the response.")
	})
	networkFirewallCmd.AddCommand(networkFirewall_listVpcEndpointAssociationsCmd)
}
