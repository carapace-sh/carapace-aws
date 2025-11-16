package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53resolver_getResolverRuleCmd = &cobra.Command{
	Use:   "get-resolver-rule",
	Short: "Gets information about a specified Resolver rule, such as the domain name that the rule forwards DNS queries for and the ID of the outbound Resolver endpoint that the rule is associated with.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53resolver_getResolverRuleCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(route53resolver_getResolverRuleCmd).Standalone()

		route53resolver_getResolverRuleCmd.Flags().String("resolver-rule-id", "", "The ID of the Resolver rule that you want to get information about.")
		route53resolver_getResolverRuleCmd.MarkFlagRequired("resolver-rule-id")
	})
	route53resolverCmd.AddCommand(route53resolver_getResolverRuleCmd)
}
