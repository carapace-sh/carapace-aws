package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var networkFirewall_createVpcEndpointAssociationCmd = &cobra.Command{
	Use:   "create-vpc-endpoint-association",
	Short: "Creates a firewall endpoint for an Network Firewall firewall.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(networkFirewall_createVpcEndpointAssociationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(networkFirewall_createVpcEndpointAssociationCmd).Standalone()

		networkFirewall_createVpcEndpointAssociationCmd.Flags().String("description", "", "A description of the VPC endpoint association.")
		networkFirewall_createVpcEndpointAssociationCmd.Flags().String("firewall-arn", "", "The Amazon Resource Name (ARN) of the firewall.")
		networkFirewall_createVpcEndpointAssociationCmd.Flags().String("subnet-mapping", "", "")
		networkFirewall_createVpcEndpointAssociationCmd.Flags().String("tags", "", "The key:value pairs to associate with the resource.")
		networkFirewall_createVpcEndpointAssociationCmd.Flags().String("vpc-id", "", "The unique identifier of the VPC where you want to create a firewall endpoint.")
		networkFirewall_createVpcEndpointAssociationCmd.MarkFlagRequired("firewall-arn")
		networkFirewall_createVpcEndpointAssociationCmd.MarkFlagRequired("subnet-mapping")
		networkFirewall_createVpcEndpointAssociationCmd.MarkFlagRequired("vpc-id")
	})
	networkFirewallCmd.AddCommand(networkFirewall_createVpcEndpointAssociationCmd)
}
