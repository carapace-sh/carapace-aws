package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53resolver_putFirewallRuleGroupPolicyCmd = &cobra.Command{
	Use:   "put-firewall-rule-group-policy",
	Short: "Attaches an Identity and Access Management (Amazon Web Services IAM) policy for sharing the rule group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53resolver_putFirewallRuleGroupPolicyCmd).Standalone()

	route53resolver_putFirewallRuleGroupPolicyCmd.Flags().String("arn", "", "The ARN (Amazon Resource Name) for the rule group that you want to share.")
	route53resolver_putFirewallRuleGroupPolicyCmd.Flags().String("firewall-rule-group-policy", "", "The Identity and Access Management (Amazon Web Services IAM) policy to attach to the rule group.")
	route53resolver_putFirewallRuleGroupPolicyCmd.MarkFlagRequired("arn")
	route53resolver_putFirewallRuleGroupPolicyCmd.MarkFlagRequired("firewall-rule-group-policy")
	route53resolverCmd.AddCommand(route53resolver_putFirewallRuleGroupPolicyCmd)
}
