package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53resolver_listResolverDnssecConfigsCmd = &cobra.Command{
	Use:   "list-resolver-dnssec-configs",
	Short: "Lists the configurations for DNSSEC validation that are associated with the current Amazon Web Services account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53resolver_listResolverDnssecConfigsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(route53resolver_listResolverDnssecConfigsCmd).Standalone()

		route53resolver_listResolverDnssecConfigsCmd.Flags().String("filters", "", "An optional specification to return a subset of objects.")
		route53resolver_listResolverDnssecConfigsCmd.Flags().String("max-results", "", "*Optional*: An integer that specifies the maximum number of DNSSEC configuration results that you want Amazon Route 53 to return.")
		route53resolver_listResolverDnssecConfigsCmd.Flags().String("next-token", "", "(Optional) If the current Amazon Web Services account has more than `MaxResults` DNSSEC configurations, use `NextToken` to get the second and subsequent pages of results.")
	})
	route53resolverCmd.AddCommand(route53resolver_listResolverDnssecConfigsCmd)
}
