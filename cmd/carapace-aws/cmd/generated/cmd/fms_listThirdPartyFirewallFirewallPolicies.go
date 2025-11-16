package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var fms_listThirdPartyFirewallFirewallPoliciesCmd = &cobra.Command{
	Use:   "list-third-party-firewall-firewall-policies",
	Short: "Retrieves a list of all of the third-party firewall policies that are associated with the third-party firewall administrator's account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(fms_listThirdPartyFirewallFirewallPoliciesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(fms_listThirdPartyFirewallFirewallPoliciesCmd).Standalone()

		fms_listThirdPartyFirewallFirewallPoliciesCmd.Flags().String("max-results", "", "The maximum number of third-party firewall policies that you want Firewall Manager to return.")
		fms_listThirdPartyFirewallFirewallPoliciesCmd.Flags().String("next-token", "", "If the previous response included a `NextToken` element, the specified third-party firewall vendor is associated with more third-party firewall policies.")
		fms_listThirdPartyFirewallFirewallPoliciesCmd.Flags().String("third-party-firewall", "", "The name of the third-party firewall vendor.")
		fms_listThirdPartyFirewallFirewallPoliciesCmd.MarkFlagRequired("max-results")
		fms_listThirdPartyFirewallFirewallPoliciesCmd.MarkFlagRequired("third-party-firewall")
	})
	fmsCmd.AddCommand(fms_listThirdPartyFirewallFirewallPoliciesCmd)
}
