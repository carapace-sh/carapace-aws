package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53resolver_associateResolverQueryLogConfigCmd = &cobra.Command{
	Use:   "associate-resolver-query-log-config",
	Short: "Associates an Amazon VPC with a specified query logging configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53resolver_associateResolverQueryLogConfigCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(route53resolver_associateResolverQueryLogConfigCmd).Standalone()

		route53resolver_associateResolverQueryLogConfigCmd.Flags().String("resolver-query-log-config-id", "", "The ID of the query logging configuration that you want to associate a VPC with.")
		route53resolver_associateResolverQueryLogConfigCmd.Flags().String("resource-id", "", "The ID of an Amazon VPC that you want this query logging configuration to log queries for.")
		route53resolver_associateResolverQueryLogConfigCmd.MarkFlagRequired("resolver-query-log-config-id")
		route53resolver_associateResolverQueryLogConfigCmd.MarkFlagRequired("resource-id")
	})
	route53resolverCmd.AddCommand(route53resolver_associateResolverQueryLogConfigCmd)
}
