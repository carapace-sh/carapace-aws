package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53resolver_updateResolverConfigCmd = &cobra.Command{
	Use:   "update-resolver-config",
	Short: "Updates the behavior configuration of Route\u00a053 Resolver behavior for a single VPC from Amazon Virtual Private Cloud.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53resolver_updateResolverConfigCmd).Standalone()

	route53resolver_updateResolverConfigCmd.Flags().String("autodefined-reverse-flag", "", "Indicates whether or not the Resolver will create autodefined rules for reverse DNS lookups.")
	route53resolver_updateResolverConfigCmd.Flags().String("resource-id", "", "The ID of the Amazon Virtual Private Cloud VPC or a Route 53 Profile that you're configuring Resolver for.")
	route53resolver_updateResolverConfigCmd.MarkFlagRequired("autodefined-reverse-flag")
	route53resolver_updateResolverConfigCmd.MarkFlagRequired("resource-id")
	route53resolverCmd.AddCommand(route53resolver_updateResolverConfigCmd)
}
