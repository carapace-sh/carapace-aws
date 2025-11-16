package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53resolver_deleteResolverEndpointCmd = &cobra.Command{
	Use:   "delete-resolver-endpoint",
	Short: "Deletes a Resolver endpoint.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53resolver_deleteResolverEndpointCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(route53resolver_deleteResolverEndpointCmd).Standalone()

		route53resolver_deleteResolverEndpointCmd.Flags().String("resolver-endpoint-id", "", "The ID of the Resolver endpoint that you want to delete.")
		route53resolver_deleteResolverEndpointCmd.MarkFlagRequired("resolver-endpoint-id")
	})
	route53resolverCmd.AddCommand(route53resolver_deleteResolverEndpointCmd)
}
