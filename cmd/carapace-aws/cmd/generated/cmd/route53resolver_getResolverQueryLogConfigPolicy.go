package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53resolver_getResolverQueryLogConfigPolicyCmd = &cobra.Command{
	Use:   "get-resolver-query-log-config-policy",
	Short: "Gets information about a query logging policy.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53resolver_getResolverQueryLogConfigPolicyCmd).Standalone()

	route53resolver_getResolverQueryLogConfigPolicyCmd.Flags().String("arn", "", "The ARN of the query logging configuration that you want to get the query logging policy for.")
	route53resolver_getResolverQueryLogConfigPolicyCmd.MarkFlagRequired("arn")
	route53resolverCmd.AddCommand(route53resolver_getResolverQueryLogConfigPolicyCmd)
}
