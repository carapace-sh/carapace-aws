package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53resolver_listResolverQueryLogConfigAssociationsCmd = &cobra.Command{
	Use:   "list-resolver-query-log-config-associations",
	Short: "Lists information about associations between Amazon VPCs and query logging configurations.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53resolver_listResolverQueryLogConfigAssociationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(route53resolver_listResolverQueryLogConfigAssociationsCmd).Standalone()

		route53resolver_listResolverQueryLogConfigAssociationsCmd.Flags().String("filters", "", "An optional specification to return a subset of query logging associations.")
		route53resolver_listResolverQueryLogConfigAssociationsCmd.Flags().String("max-results", "", "The maximum number of query logging associations that you want to return in the response to a `ListResolverQueryLogConfigAssociations` request.")
		route53resolver_listResolverQueryLogConfigAssociationsCmd.Flags().String("next-token", "", "For the first `ListResolverQueryLogConfigAssociations` request, omit this value.")
		route53resolver_listResolverQueryLogConfigAssociationsCmd.Flags().String("sort-by", "", "The element that you want Resolver to sort query logging associations by.")
		route53resolver_listResolverQueryLogConfigAssociationsCmd.Flags().String("sort-order", "", "If you specified a value for `SortBy`, the order that you want query logging associations to be listed in, `ASCENDING` or `DESCENDING`.")
	})
	route53resolverCmd.AddCommand(route53resolver_listResolverQueryLogConfigAssociationsCmd)
}
