package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53resolver_deleteFirewallRuleGroupCmd = &cobra.Command{
	Use:   "delete-firewall-rule-group",
	Short: "Deletes the specified firewall rule group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53resolver_deleteFirewallRuleGroupCmd).Standalone()

	route53resolver_deleteFirewallRuleGroupCmd.Flags().String("firewall-rule-group-id", "", "The unique identifier of the firewall rule group that you want to delete.")
	route53resolver_deleteFirewallRuleGroupCmd.MarkFlagRequired("firewall-rule-group-id")
	route53resolverCmd.AddCommand(route53resolver_deleteFirewallRuleGroupCmd)
}
