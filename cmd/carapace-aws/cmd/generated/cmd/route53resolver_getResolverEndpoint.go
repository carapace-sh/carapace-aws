package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53resolver_getResolverEndpointCmd = &cobra.Command{
	Use:   "get-resolver-endpoint",
	Short: "Gets information about a specified Resolver endpoint, such as whether it's an inbound or an outbound Resolver endpoint, and the current status of the endpoint.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53resolver_getResolverEndpointCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(route53resolver_getResolverEndpointCmd).Standalone()

		route53resolver_getResolverEndpointCmd.Flags().String("resolver-endpoint-id", "", "The ID of the Resolver endpoint that you want to get information about.")
		route53resolver_getResolverEndpointCmd.MarkFlagRequired("resolver-endpoint-id")
	})
	route53resolverCmd.AddCommand(route53resolver_getResolverEndpointCmd)
}
