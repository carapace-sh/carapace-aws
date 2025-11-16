package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53resolver_listFirewallRulesCmd = &cobra.Command{
	Use:   "list-firewall-rules",
	Short: "Retrieves the firewall rules that you have defined for the specified firewall rule group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53resolver_listFirewallRulesCmd).Standalone()

	route53resolver_listFirewallRulesCmd.Flags().String("action", "", "Optional additional filter for the rules to retrieve.")
	route53resolver_listFirewallRulesCmd.Flags().String("firewall-rule-group-id", "", "The unique identifier of the firewall rule group that you want to retrieve the rules for.")
	route53resolver_listFirewallRulesCmd.Flags().String("max-results", "", "The maximum number of objects that you want Resolver to return for this request.")
	route53resolver_listFirewallRulesCmd.Flags().String("next-token", "", "For the first call to this list request, omit this value.")
	route53resolver_listFirewallRulesCmd.Flags().String("priority", "", "Optional additional filter for the rules to retrieve.")
	route53resolver_listFirewallRulesCmd.MarkFlagRequired("firewall-rule-group-id")
	route53resolverCmd.AddCommand(route53resolver_listFirewallRulesCmd)
}
