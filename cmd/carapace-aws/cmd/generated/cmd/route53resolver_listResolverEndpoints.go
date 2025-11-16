package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53resolver_listResolverEndpointsCmd = &cobra.Command{
	Use:   "list-resolver-endpoints",
	Short: "Lists all the Resolver endpoints that were created using the current Amazon Web Services account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53resolver_listResolverEndpointsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(route53resolver_listResolverEndpointsCmd).Standalone()

		route53resolver_listResolverEndpointsCmd.Flags().String("filters", "", "An optional specification to return a subset of Resolver endpoints, such as all inbound Resolver endpoints.")
		route53resolver_listResolverEndpointsCmd.Flags().String("max-results", "", "The maximum number of Resolver endpoints that you want to return in the response to a `ListResolverEndpoints` request.")
		route53resolver_listResolverEndpointsCmd.Flags().String("next-token", "", "For the first `ListResolverEndpoints` request, omit this value.")
	})
	route53resolverCmd.AddCommand(route53resolver_listResolverEndpointsCmd)
}
