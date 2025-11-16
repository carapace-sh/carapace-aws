package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53resolver_getFirewallRuleGroupPolicyCmd = &cobra.Command{
	Use:   "get-firewall-rule-group-policy",
	Short: "Returns the Identity and Access Management (Amazon Web Services IAM) policy for sharing the specified rule group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53resolver_getFirewallRuleGroupPolicyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(route53resolver_getFirewallRuleGroupPolicyCmd).Standalone()

		route53resolver_getFirewallRuleGroupPolicyCmd.Flags().String("arn", "", "The ARN (Amazon Resource Name) for the rule group.")
		route53resolver_getFirewallRuleGroupPolicyCmd.MarkFlagRequired("arn")
	})
	route53resolverCmd.AddCommand(route53resolver_getFirewallRuleGroupPolicyCmd)
}
