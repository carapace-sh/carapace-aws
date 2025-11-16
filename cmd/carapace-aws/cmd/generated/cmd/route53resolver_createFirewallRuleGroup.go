package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53resolver_createFirewallRuleGroupCmd = &cobra.Command{
	Use:   "create-firewall-rule-group",
	Short: "Creates an empty DNS Firewall rule group for filtering DNS network traffic in a VPC.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53resolver_createFirewallRuleGroupCmd).Standalone()

	route53resolver_createFirewallRuleGroupCmd.Flags().String("creator-request-id", "", "A unique string defined by you to identify the request.")
	route53resolver_createFirewallRuleGroupCmd.Flags().String("name", "", "A name that lets you identify the rule group, to manage and use it.")
	route53resolver_createFirewallRuleGroupCmd.Flags().String("tags", "", "A list of the tag keys and values that you want to associate with the rule group.")
	route53resolver_createFirewallRuleGroupCmd.MarkFlagRequired("creator-request-id")
	route53resolver_createFirewallRuleGroupCmd.MarkFlagRequired("name")
	route53resolverCmd.AddCommand(route53resolver_createFirewallRuleGroupCmd)
}
