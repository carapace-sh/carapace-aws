package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var networkFirewall_listRuleGroupsCmd = &cobra.Command{
	Use:   "list-rule-groups",
	Short: "Retrieves the metadata for the rule groups that you have defined.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(networkFirewall_listRuleGroupsCmd).Standalone()

	networkFirewall_listRuleGroupsCmd.Flags().String("managed-type", "", "Indicates the general category of the Amazon Web Services managed rule group.")
	networkFirewall_listRuleGroupsCmd.Flags().String("max-results", "", "The maximum number of objects that you want Network Firewall to return for this request.")
	networkFirewall_listRuleGroupsCmd.Flags().String("next-token", "", "When you request a list of objects with a `MaxResults` setting, if the number of objects that are still available for retrieval exceeds the maximum you requested, Network Firewall returns a `NextToken` value in the response.")
	networkFirewall_listRuleGroupsCmd.Flags().String("scope", "", "The scope of the request.")
	networkFirewall_listRuleGroupsCmd.Flags().String("type", "", "Indicates whether the rule group is stateless or stateful.")
	networkFirewallCmd.AddCommand(networkFirewall_listRuleGroupsCmd)
}
