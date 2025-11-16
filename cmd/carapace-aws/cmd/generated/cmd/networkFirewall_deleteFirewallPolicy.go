package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var networkFirewall_deleteFirewallPolicyCmd = &cobra.Command{
	Use:   "delete-firewall-policy",
	Short: "Deletes the specified [FirewallPolicy]().",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(networkFirewall_deleteFirewallPolicyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(networkFirewall_deleteFirewallPolicyCmd).Standalone()

		networkFirewall_deleteFirewallPolicyCmd.Flags().String("firewall-policy-arn", "", "The Amazon Resource Name (ARN) of the firewall policy.")
		networkFirewall_deleteFirewallPolicyCmd.Flags().String("firewall-policy-name", "", "The descriptive name of the firewall policy.")
	})
	networkFirewallCmd.AddCommand(networkFirewall_deleteFirewallPolicyCmd)
}
