package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53resolver_getResolverQueryLogConfigAssociationCmd = &cobra.Command{
	Use:   "get-resolver-query-log-config-association",
	Short: "Gets information about a specified association between a Resolver query logging configuration and an Amazon VPC.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53resolver_getResolverQueryLogConfigAssociationCmd).Standalone()

	route53resolver_getResolverQueryLogConfigAssociationCmd.Flags().String("resolver-query-log-config-association-id", "", "The ID of the Resolver query logging configuration association that you want to get information about.")
	route53resolver_getResolverQueryLogConfigAssociationCmd.MarkFlagRequired("resolver-query-log-config-association-id")
	route53resolverCmd.AddCommand(route53resolver_getResolverQueryLogConfigAssociationCmd)
}
