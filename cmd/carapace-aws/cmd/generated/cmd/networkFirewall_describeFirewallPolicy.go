package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var networkFirewall_describeFirewallPolicyCmd = &cobra.Command{
	Use:   "describe-firewall-policy",
	Short: "Returns the data objects for the specified firewall policy.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(networkFirewall_describeFirewallPolicyCmd).Standalone()

	networkFirewall_describeFirewallPolicyCmd.Flags().String("firewall-policy-arn", "", "The Amazon Resource Name (ARN) of the firewall policy.")
	networkFirewall_describeFirewallPolicyCmd.Flags().String("firewall-policy-name", "", "The descriptive name of the firewall policy.")
	networkFirewallCmd.AddCommand(networkFirewall_describeFirewallPolicyCmd)
}
