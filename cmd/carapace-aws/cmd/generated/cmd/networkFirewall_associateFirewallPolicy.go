package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var networkFirewall_associateFirewallPolicyCmd = &cobra.Command{
	Use:   "associate-firewall-policy",
	Short: "Associates a [FirewallPolicy]() to a [Firewall]().",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(networkFirewall_associateFirewallPolicyCmd).Standalone()

	networkFirewall_associateFirewallPolicyCmd.Flags().String("firewall-arn", "", "The Amazon Resource Name (ARN) of the firewall.")
	networkFirewall_associateFirewallPolicyCmd.Flags().String("firewall-name", "", "The descriptive name of the firewall.")
	networkFirewall_associateFirewallPolicyCmd.Flags().String("firewall-policy-arn", "", "The Amazon Resource Name (ARN) of the firewall policy.")
	networkFirewall_associateFirewallPolicyCmd.Flags().String("update-token", "", "An optional token that you can use for optimistic locking.")
	networkFirewall_associateFirewallPolicyCmd.MarkFlagRequired("firewall-policy-arn")
	networkFirewallCmd.AddCommand(networkFirewall_associateFirewallPolicyCmd)
}
