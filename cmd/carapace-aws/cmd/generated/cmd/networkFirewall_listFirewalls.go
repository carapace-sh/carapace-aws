package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var networkFirewall_listFirewallsCmd = &cobra.Command{
	Use:   "list-firewalls",
	Short: "Retrieves the metadata for the firewalls that you have defined.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(networkFirewall_listFirewallsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(networkFirewall_listFirewallsCmd).Standalone()

		networkFirewall_listFirewallsCmd.Flags().String("max-results", "", "The maximum number of objects that you want Network Firewall to return for this request.")
		networkFirewall_listFirewallsCmd.Flags().String("next-token", "", "When you request a list of objects with a `MaxResults` setting, if the number of objects that are still available for retrieval exceeds the maximum you requested, Network Firewall returns a `NextToken` value in the response.")
		networkFirewall_listFirewallsCmd.Flags().String("vpc-ids", "", "The unique identifiers of the VPCs that you want Network Firewall to retrieve the firewalls for.")
	})
	networkFirewallCmd.AddCommand(networkFirewall_listFirewallsCmd)
}
