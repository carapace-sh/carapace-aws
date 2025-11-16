package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53resolver_listFirewallRuleGroupAssociationsCmd = &cobra.Command{
	Use:   "list-firewall-rule-group-associations",
	Short: "Retrieves the firewall rule group associations that you have defined.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53resolver_listFirewallRuleGroupAssociationsCmd).Standalone()

	route53resolver_listFirewallRuleGroupAssociationsCmd.Flags().String("firewall-rule-group-id", "", "The unique identifier of the firewall rule group that you want to retrieve the associations for.")
	route53resolver_listFirewallRuleGroupAssociationsCmd.Flags().String("max-results", "", "The maximum number of objects that you want Resolver to return for this request.")
	route53resolver_listFirewallRuleGroupAssociationsCmd.Flags().String("next-token", "", "For the first call to this list request, omit this value.")
	route53resolver_listFirewallRuleGroupAssociationsCmd.Flags().String("priority", "", "The setting that determines the processing order of the rule group among the rule groups that are associated with a single VPC.")
	route53resolver_listFirewallRuleGroupAssociationsCmd.Flags().String("status", "", "The association `Status` setting that you want DNS Firewall to filter on for the list.")
	route53resolver_listFirewallRuleGroupAssociationsCmd.Flags().String("vpc-id", "", "The unique identifier of the VPC that you want to retrieve the associations for.")
	route53resolverCmd.AddCommand(route53resolver_listFirewallRuleGroupAssociationsCmd)
}
