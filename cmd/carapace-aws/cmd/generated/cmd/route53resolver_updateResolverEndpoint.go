package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53resolver_updateResolverEndpointCmd = &cobra.Command{
	Use:   "update-resolver-endpoint",
	Short: "Updates the name, or endpoint type for an inbound or an outbound Resolver endpoint.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53resolver_updateResolverEndpointCmd).Standalone()

	route53resolver_updateResolverEndpointCmd.Flags().String("name", "", "The name of the Resolver endpoint that you want to update.")
	route53resolver_updateResolverEndpointCmd.Flags().String("protocols", "", "The protocols you want to use for the endpoint.")
	route53resolver_updateResolverEndpointCmd.Flags().String("resolver-endpoint-id", "", "The ID of the Resolver endpoint that you want to update.")
	route53resolver_updateResolverEndpointCmd.Flags().String("resolver-endpoint-type", "", "Specifies the endpoint type for what type of IP address the endpoint uses to forward DNS queries.")
	route53resolver_updateResolverEndpointCmd.Flags().String("update-ip-addresses", "", "Specifies the IPv6 address when you update the Resolver endpoint from IPv4 to dual-stack.")
	route53resolver_updateResolverEndpointCmd.MarkFlagRequired("resolver-endpoint-id")
	route53resolverCmd.AddCommand(route53resolver_updateResolverEndpointCmd)
}
