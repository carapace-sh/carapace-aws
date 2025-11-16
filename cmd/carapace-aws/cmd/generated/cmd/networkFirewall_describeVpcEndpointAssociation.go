package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var networkFirewall_describeVpcEndpointAssociationCmd = &cobra.Command{
	Use:   "describe-vpc-endpoint-association",
	Short: "Returns the data object for the specified VPC endpoint association.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(networkFirewall_describeVpcEndpointAssociationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(networkFirewall_describeVpcEndpointAssociationCmd).Standalone()

		networkFirewall_describeVpcEndpointAssociationCmd.Flags().String("vpc-endpoint-association-arn", "", "The Amazon Resource Name (ARN) of a VPC endpoint association.")
		networkFirewall_describeVpcEndpointAssociationCmd.MarkFlagRequired("vpc-endpoint-association-arn")
	})
	networkFirewallCmd.AddCommand(networkFirewall_describeVpcEndpointAssociationCmd)
}
