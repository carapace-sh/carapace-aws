package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53resolver_createResolverRuleCmd = &cobra.Command{
	Use:   "create-resolver-rule",
	Short: "For DNS queries that originate in your VPCs, specifies which Resolver endpoint the queries pass through, one domain name that you want to forward to your network, and the IP addresses of the DNS resolvers in your network.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53resolver_createResolverRuleCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(route53resolver_createResolverRuleCmd).Standalone()

		route53resolver_createResolverRuleCmd.Flags().String("creator-request-id", "", "A unique string that identifies the request and that allows failed requests to be retried without the risk of running the operation twice.")
		route53resolver_createResolverRuleCmd.Flags().String("delegation-record", "", "DNS queries with the delegation records that match this domain name are forwarded to the resolvers on your network.")
		route53resolver_createResolverRuleCmd.Flags().String("domain-name", "", "DNS queries for this domain name are forwarded to the IP addresses that you specify in `TargetIps`.")
		route53resolver_createResolverRuleCmd.Flags().String("name", "", "A friendly name that lets you easily find a rule in the Resolver dashboard in the Route 53 console.")
		route53resolver_createResolverRuleCmd.Flags().String("resolver-endpoint-id", "", "The ID of the outbound Resolver endpoint that you want to use to route DNS queries to the IP addresses that you specify in `TargetIps`.")
		route53resolver_createResolverRuleCmd.Flags().String("rule-type", "", "When you want to forward DNS queries for specified domain name to resolvers on your network, specify `FORWARD` or `DELEGATE`.")
		route53resolver_createResolverRuleCmd.Flags().String("tags", "", "A list of the tag keys and values that you want to associate with the endpoint.")
		route53resolver_createResolverRuleCmd.Flags().String("target-ips", "", "The IPs that you want Resolver to forward DNS queries to.")
		route53resolver_createResolverRuleCmd.MarkFlagRequired("creator-request-id")
		route53resolver_createResolverRuleCmd.MarkFlagRequired("rule-type")
	})
	route53resolverCmd.AddCommand(route53resolver_createResolverRuleCmd)
}
