package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53resolver_putResolverRulePolicyCmd = &cobra.Command{
	Use:   "put-resolver-rule-policy",
	Short: "Specifies an Amazon Web Services rule that you want to share with another account, the account that you want to share the rule with, and the operations that you want the account to be able to perform on the rule.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53resolver_putResolverRulePolicyCmd).Standalone()

	route53resolver_putResolverRulePolicyCmd.Flags().String("arn", "", "The Amazon Resource Name (ARN) of the rule that you want to share with another account.")
	route53resolver_putResolverRulePolicyCmd.Flags().String("resolver-rule-policy", "", "An Identity and Access Management policy statement that lists the rules that you want to share with another Amazon Web Services account and the operations that you want the account to be able to perform.")
	route53resolver_putResolverRulePolicyCmd.MarkFlagRequired("arn")
	route53resolver_putResolverRulePolicyCmd.MarkFlagRequired("resolver-rule-policy")
	route53resolverCmd.AddCommand(route53resolver_putResolverRulePolicyCmd)
}
