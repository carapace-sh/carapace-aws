package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53resolver_getResolverRuleAssociationCmd = &cobra.Command{
	Use:   "get-resolver-rule-association",
	Short: "Gets information about an association between a specified Resolver rule and a VPC.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53resolver_getResolverRuleAssociationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(route53resolver_getResolverRuleAssociationCmd).Standalone()

		route53resolver_getResolverRuleAssociationCmd.Flags().String("resolver-rule-association-id", "", "The ID of the Resolver rule association that you want to get information about.")
		route53resolver_getResolverRuleAssociationCmd.MarkFlagRequired("resolver-rule-association-id")
	})
	route53resolverCmd.AddCommand(route53resolver_getResolverRuleAssociationCmd)
}
