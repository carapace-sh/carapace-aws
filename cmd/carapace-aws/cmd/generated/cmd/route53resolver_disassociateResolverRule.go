package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53resolver_disassociateResolverRuleCmd = &cobra.Command{
	Use:   "disassociate-resolver-rule",
	Short: "Removes the association between a specified Resolver rule and a specified VPC.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53resolver_disassociateResolverRuleCmd).Standalone()

	route53resolver_disassociateResolverRuleCmd.Flags().String("resolver-rule-id", "", "The ID of the Resolver rule that you want to disassociate from the specified VPC.")
	route53resolver_disassociateResolverRuleCmd.Flags().String("vpcid", "", "The ID of the VPC that you want to disassociate the Resolver rule from.")
	route53resolver_disassociateResolverRuleCmd.MarkFlagRequired("resolver-rule-id")
	route53resolver_disassociateResolverRuleCmd.MarkFlagRequired("vpcid")
	route53resolverCmd.AddCommand(route53resolver_disassociateResolverRuleCmd)
}
