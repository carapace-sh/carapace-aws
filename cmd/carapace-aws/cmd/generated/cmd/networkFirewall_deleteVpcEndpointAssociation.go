package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var networkFirewall_deleteVpcEndpointAssociationCmd = &cobra.Command{
	Use:   "delete-vpc-endpoint-association",
	Short: "Deletes the specified [VpcEndpointAssociation]().",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(networkFirewall_deleteVpcEndpointAssociationCmd).Standalone()

	networkFirewall_deleteVpcEndpointAssociationCmd.Flags().String("vpc-endpoint-association-arn", "", "The Amazon Resource Name (ARN) of a VPC endpoint association.")
	networkFirewall_deleteVpcEndpointAssociationCmd.MarkFlagRequired("vpc-endpoint-association-arn")
	networkFirewallCmd.AddCommand(networkFirewall_deleteVpcEndpointAssociationCmd)
}
