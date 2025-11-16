package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53resolver_getResolverQueryLogConfigCmd = &cobra.Command{
	Use:   "get-resolver-query-log-config",
	Short: "Gets information about a specified Resolver query logging configuration, such as the number of VPCs that the configuration is logging queries for and the location that logs are sent to.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53resolver_getResolverQueryLogConfigCmd).Standalone()

	route53resolver_getResolverQueryLogConfigCmd.Flags().String("resolver-query-log-config-id", "", "The ID of the Resolver query logging configuration that you want to get information about.")
	route53resolver_getResolverQueryLogConfigCmd.MarkFlagRequired("resolver-query-log-config-id")
	route53resolverCmd.AddCommand(route53resolver_getResolverQueryLogConfigCmd)
}
