package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53resolver_listResolverEndpointIpAddressesCmd = &cobra.Command{
	Use:   "list-resolver-endpoint-ip-addresses",
	Short: "Gets the IP addresses for a specified Resolver endpoint.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53resolver_listResolverEndpointIpAddressesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(route53resolver_listResolverEndpointIpAddressesCmd).Standalone()

		route53resolver_listResolverEndpointIpAddressesCmd.Flags().String("max-results", "", "The maximum number of IP addresses that you want to return in the response to a `ListResolverEndpointIpAddresses` request.")
		route53resolver_listResolverEndpointIpAddressesCmd.Flags().String("next-token", "", "For the first `ListResolverEndpointIpAddresses` request, omit this value.")
		route53resolver_listResolverEndpointIpAddressesCmd.Flags().String("resolver-endpoint-id", "", "The ID of the Resolver endpoint that you want to get IP addresses for.")
		route53resolver_listResolverEndpointIpAddressesCmd.MarkFlagRequired("resolver-endpoint-id")
	})
	route53resolverCmd.AddCommand(route53resolver_listResolverEndpointIpAddressesCmd)
}
