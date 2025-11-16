package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53resolver_getResolverConfigCmd = &cobra.Command{
	Use:   "get-resolver-config",
	Short: "Retrieves the behavior configuration of Route\u00a053 Resolver behavior for a single VPC from Amazon Virtual Private Cloud.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53resolver_getResolverConfigCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(route53resolver_getResolverConfigCmd).Standalone()

		route53resolver_getResolverConfigCmd.Flags().String("resource-id", "", "Resource ID of the Amazon VPC that you want to get information about.")
		route53resolver_getResolverConfigCmd.MarkFlagRequired("resource-id")
	})
	route53resolverCmd.AddCommand(route53resolver_getResolverConfigCmd)
}
