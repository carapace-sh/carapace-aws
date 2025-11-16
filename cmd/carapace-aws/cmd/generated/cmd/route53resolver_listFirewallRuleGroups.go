package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53resolver_listFirewallRuleGroupsCmd = &cobra.Command{
	Use:   "list-firewall-rule-groups",
	Short: "Retrieves the minimal high-level information for the rule groups that you have defined.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53resolver_listFirewallRuleGroupsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(route53resolver_listFirewallRuleGroupsCmd).Standalone()

		route53resolver_listFirewallRuleGroupsCmd.Flags().String("max-results", "", "The maximum number of objects that you want Resolver to return for this request.")
		route53resolver_listFirewallRuleGroupsCmd.Flags().String("next-token", "", "For the first call to this list request, omit this value.")
	})
	route53resolverCmd.AddCommand(route53resolver_listFirewallRuleGroupsCmd)
}
