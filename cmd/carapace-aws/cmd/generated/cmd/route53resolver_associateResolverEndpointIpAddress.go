package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53resolver_associateResolverEndpointIpAddressCmd = &cobra.Command{
	Use:   "associate-resolver-endpoint-ip-address",
	Short: "Adds IP addresses to an inbound or an outbound Resolver endpoint.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53resolver_associateResolverEndpointIpAddressCmd).Standalone()

	route53resolver_associateResolverEndpointIpAddressCmd.Flags().String("ip-address", "", "Either the IPv4 address that you want to add to a Resolver endpoint or a subnet ID.")
	route53resolver_associateResolverEndpointIpAddressCmd.Flags().String("resolver-endpoint-id", "", "The ID of the Resolver endpoint that you want to associate IP addresses with.")
	route53resolver_associateResolverEndpointIpAddressCmd.MarkFlagRequired("ip-address")
	route53resolver_associateResolverEndpointIpAddressCmd.MarkFlagRequired("resolver-endpoint-id")
	route53resolverCmd.AddCommand(route53resolver_associateResolverEndpointIpAddressCmd)
}
