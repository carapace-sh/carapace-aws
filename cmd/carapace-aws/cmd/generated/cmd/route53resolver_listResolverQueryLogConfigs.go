package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53resolver_listResolverQueryLogConfigsCmd = &cobra.Command{
	Use:   "list-resolver-query-log-configs",
	Short: "Lists information about the specified query logging configurations.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53resolver_listResolverQueryLogConfigsCmd).Standalone()

	route53resolver_listResolverQueryLogConfigsCmd.Flags().String("filters", "", "An optional specification to return a subset of query logging configurations.")
	route53resolver_listResolverQueryLogConfigsCmd.Flags().String("max-results", "", "The maximum number of query logging configurations that you want to return in the response to a `ListResolverQueryLogConfigs` request.")
	route53resolver_listResolverQueryLogConfigsCmd.Flags().String("next-token", "", "For the first `ListResolverQueryLogConfigs` request, omit this value.")
	route53resolver_listResolverQueryLogConfigsCmd.Flags().String("sort-by", "", "The element that you want Resolver to sort query logging configurations by.")
	route53resolver_listResolverQueryLogConfigsCmd.Flags().String("sort-order", "", "If you specified a value for `SortBy`, the order that you want query logging configurations to be listed in, `ASCENDING` or `DESCENDING`.")
	route53resolverCmd.AddCommand(route53resolver_listResolverQueryLogConfigsCmd)
}
