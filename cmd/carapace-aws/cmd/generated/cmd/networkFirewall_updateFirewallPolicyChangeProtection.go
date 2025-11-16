package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var networkFirewall_updateFirewallPolicyChangeProtectionCmd = &cobra.Command{
	Use:   "update-firewall-policy-change-protection",
	Short: "Modifies the flag, `ChangeProtection`, which indicates whether it is possible to change the firewall.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(networkFirewall_updateFirewallPolicyChangeProtectionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(networkFirewall_updateFirewallPolicyChangeProtectionCmd).Standalone()

		networkFirewall_updateFirewallPolicyChangeProtectionCmd.Flags().String("firewall-arn", "", "The Amazon Resource Name (ARN) of the firewall.")
		networkFirewall_updateFirewallPolicyChangeProtectionCmd.Flags().String("firewall-name", "", "The descriptive name of the firewall.")
		networkFirewall_updateFirewallPolicyChangeProtectionCmd.Flags().Bool("firewall-policy-change-protection", false, "A setting indicating whether the firewall is protected against a change to the firewall policy association.")
		networkFirewall_updateFirewallPolicyChangeProtectionCmd.Flags().Bool("no-firewall-policy-change-protection", false, "A setting indicating whether the firewall is protected against a change to the firewall policy association.")
		networkFirewall_updateFirewallPolicyChangeProtectionCmd.Flags().String("update-token", "", "An optional token that you can use for optimistic locking.")
		networkFirewall_updateFirewallPolicyChangeProtectionCmd.MarkFlagRequired("firewall-policy-change-protection")
		networkFirewall_updateFirewallPolicyChangeProtectionCmd.Flag("no-firewall-policy-change-protection").Hidden = true
		networkFirewall_updateFirewallPolicyChangeProtectionCmd.MarkFlagRequired("no-firewall-policy-change-protection")
	})
	networkFirewallCmd.AddCommand(networkFirewall_updateFirewallPolicyChangeProtectionCmd)
}
