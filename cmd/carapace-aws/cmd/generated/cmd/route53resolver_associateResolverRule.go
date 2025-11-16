package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53resolver_associateResolverRuleCmd = &cobra.Command{
	Use:   "associate-resolver-rule",
	Short: "Associates a Resolver rule with a VPC.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53resolver_associateResolverRuleCmd).Standalone()

	route53resolver_associateResolverRuleCmd.Flags().String("name", "", "A name for the association that you're creating between a Resolver rule and a VPC.")
	route53resolver_associateResolverRuleCmd.Flags().String("resolver-rule-id", "", "The ID of the Resolver rule that you want to associate with the VPC.")
	route53resolver_associateResolverRuleCmd.Flags().String("vpcid", "", "The ID of the VPC that you want to associate the Resolver rule with.")
	route53resolver_associateResolverRuleCmd.MarkFlagRequired("resolver-rule-id")
	route53resolver_associateResolverRuleCmd.MarkFlagRequired("vpcid")
	route53resolverCmd.AddCommand(route53resolver_associateResolverRuleCmd)
}
