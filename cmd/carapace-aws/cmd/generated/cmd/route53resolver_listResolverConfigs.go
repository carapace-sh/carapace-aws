package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53resolver_listResolverConfigsCmd = &cobra.Command{
	Use:   "list-resolver-configs",
	Short: "Retrieves the Resolver configurations that you have defined.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53resolver_listResolverConfigsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(route53resolver_listResolverConfigsCmd).Standalone()

		route53resolver_listResolverConfigsCmd.Flags().String("max-results", "", "The maximum number of Resolver configurations that you want to return in the response to a `ListResolverConfigs` request.")
		route53resolver_listResolverConfigsCmd.Flags().String("next-token", "", "(Optional) If the current Amazon Web Services account has more than `MaxResults` Resolver configurations, use `NextToken` to get the second and subsequent pages of results.")
	})
	route53resolverCmd.AddCommand(route53resolver_listResolverConfigsCmd)
}
