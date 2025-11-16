package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53resolver_disassociateResolverQueryLogConfigCmd = &cobra.Command{
	Use:   "disassociate-resolver-query-log-config",
	Short: "Disassociates a VPC from a query logging configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53resolver_disassociateResolverQueryLogConfigCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(route53resolver_disassociateResolverQueryLogConfigCmd).Standalone()

		route53resolver_disassociateResolverQueryLogConfigCmd.Flags().String("resolver-query-log-config-id", "", "The ID of the query logging configuration that you want to disassociate a specified VPC from.")
		route53resolver_disassociateResolverQueryLogConfigCmd.Flags().String("resource-id", "", "The ID of the Amazon VPC that you want to disassociate from a specified query logging configuration.")
		route53resolver_disassociateResolverQueryLogConfigCmd.MarkFlagRequired("resolver-query-log-config-id")
		route53resolver_disassociateResolverQueryLogConfigCmd.MarkFlagRequired("resource-id")
	})
	route53resolverCmd.AddCommand(route53resolver_disassociateResolverQueryLogConfigCmd)
}
