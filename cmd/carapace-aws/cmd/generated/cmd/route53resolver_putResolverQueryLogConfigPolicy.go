package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53resolver_putResolverQueryLogConfigPolicyCmd = &cobra.Command{
	Use:   "put-resolver-query-log-config-policy",
	Short: "Specifies an Amazon Web Services account that you want to share a query logging configuration with, the query logging configuration that you want to share, and the operations that you want the account to be able to perform on the configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53resolver_putResolverQueryLogConfigPolicyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(route53resolver_putResolverQueryLogConfigPolicyCmd).Standalone()

		route53resolver_putResolverQueryLogConfigPolicyCmd.Flags().String("arn", "", "The Amazon Resource Name (ARN) of the account that you want to share rules with.")
		route53resolver_putResolverQueryLogConfigPolicyCmd.Flags().String("resolver-query-log-config-policy", "", "An Identity and Access Management policy statement that lists the query logging configurations that you want to share with another Amazon Web Services account and the operations that you want the account to be able to perform.")
		route53resolver_putResolverQueryLogConfigPolicyCmd.MarkFlagRequired("arn")
		route53resolver_putResolverQueryLogConfigPolicyCmd.MarkFlagRequired("resolver-query-log-config-policy")
	})
	route53resolverCmd.AddCommand(route53resolver_putResolverQueryLogConfigPolicyCmd)
}
