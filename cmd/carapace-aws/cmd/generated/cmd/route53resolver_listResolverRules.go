package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53resolver_listResolverRulesCmd = &cobra.Command{
	Use:   "list-resolver-rules",
	Short: "Lists the Resolver rules that were created using the current Amazon Web Services account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53resolver_listResolverRulesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(route53resolver_listResolverRulesCmd).Standalone()

		route53resolver_listResolverRulesCmd.Flags().String("filters", "", "An optional specification to return a subset of Resolver rules, such as all Resolver rules that are associated with the same Resolver endpoint.")
		route53resolver_listResolverRulesCmd.Flags().String("max-results", "", "The maximum number of Resolver rules that you want to return in the response to a `ListResolverRules` request.")
		route53resolver_listResolverRulesCmd.Flags().String("next-token", "", "For the first `ListResolverRules` request, omit this value.")
	})
	route53resolverCmd.AddCommand(route53resolver_listResolverRulesCmd)
}
