package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var networkFirewall_listFirewallPoliciesCmd = &cobra.Command{
	Use:   "list-firewall-policies",
	Short: "Retrieves the metadata for the firewall policies that you have defined.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(networkFirewall_listFirewallPoliciesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(networkFirewall_listFirewallPoliciesCmd).Standalone()

		networkFirewall_listFirewallPoliciesCmd.Flags().String("max-results", "", "The maximum number of objects that you want Network Firewall to return for this request.")
		networkFirewall_listFirewallPoliciesCmd.Flags().String("next-token", "", "When you request a list of objects with a `MaxResults` setting, if the number of objects that are still available for retrieval exceeds the maximum you requested, Network Firewall returns a `NextToken` value in the response.")
	})
	networkFirewallCmd.AddCommand(networkFirewall_listFirewallPoliciesCmd)
}
