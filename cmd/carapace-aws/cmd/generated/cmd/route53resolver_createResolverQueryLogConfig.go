package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53resolver_createResolverQueryLogConfigCmd = &cobra.Command{
	Use:   "create-resolver-query-log-config",
	Short: "Creates a Resolver query logging configuration, which defines where you want Resolver to save DNS query logs that originate in your VPCs.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53resolver_createResolverQueryLogConfigCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(route53resolver_createResolverQueryLogConfigCmd).Standalone()

		route53resolver_createResolverQueryLogConfigCmd.Flags().String("creator-request-id", "", "A unique string that identifies the request and that allows failed requests to be retried without the risk of running the operation twice.")
		route53resolver_createResolverQueryLogConfigCmd.Flags().String("destination-arn", "", "The ARN of the resource that you want Resolver to send query logs.")
		route53resolver_createResolverQueryLogConfigCmd.Flags().String("name", "", "The name that you want to give the query logging configuration.")
		route53resolver_createResolverQueryLogConfigCmd.Flags().String("tags", "", "A list of the tag keys and values that you want to associate with the query logging configuration.")
		route53resolver_createResolverQueryLogConfigCmd.MarkFlagRequired("creator-request-id")
		route53resolver_createResolverQueryLogConfigCmd.MarkFlagRequired("destination-arn")
		route53resolver_createResolverQueryLogConfigCmd.MarkFlagRequired("name")
	})
	route53resolverCmd.AddCommand(route53resolver_createResolverQueryLogConfigCmd)
}
