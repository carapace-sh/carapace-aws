package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53resolver_disassociateResolverEndpointIpAddressCmd = &cobra.Command{
	Use:   "disassociate-resolver-endpoint-ip-address",
	Short: "Removes IP addresses from an inbound or an outbound Resolver endpoint.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53resolver_disassociateResolverEndpointIpAddressCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(route53resolver_disassociateResolverEndpointIpAddressCmd).Standalone()

		route53resolver_disassociateResolverEndpointIpAddressCmd.Flags().String("ip-address", "", "The IPv4 address that you want to remove from a Resolver endpoint.")
		route53resolver_disassociateResolverEndpointIpAddressCmd.Flags().String("resolver-endpoint-id", "", "The ID of the Resolver endpoint that you want to disassociate an IP address from.")
		route53resolver_disassociateResolverEndpointIpAddressCmd.MarkFlagRequired("ip-address")
		route53resolver_disassociateResolverEndpointIpAddressCmd.MarkFlagRequired("resolver-endpoint-id")
	})
	route53resolverCmd.AddCommand(route53resolver_disassociateResolverEndpointIpAddressCmd)
}
