package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var networkFirewall_updateFirewallDeleteProtectionCmd = &cobra.Command{
	Use:   "update-firewall-delete-protection",
	Short: "Modifies the flag, `DeleteProtection`, which indicates whether it is possible to delete the firewall.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(networkFirewall_updateFirewallDeleteProtectionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(networkFirewall_updateFirewallDeleteProtectionCmd).Standalone()

		networkFirewall_updateFirewallDeleteProtectionCmd.Flags().Bool("delete-protection", false, "A flag indicating whether it is possible to delete the firewall.")
		networkFirewall_updateFirewallDeleteProtectionCmd.Flags().String("firewall-arn", "", "The Amazon Resource Name (ARN) of the firewall.")
		networkFirewall_updateFirewallDeleteProtectionCmd.Flags().String("firewall-name", "", "The descriptive name of the firewall.")
		networkFirewall_updateFirewallDeleteProtectionCmd.Flags().Bool("no-delete-protection", false, "A flag indicating whether it is possible to delete the firewall.")
		networkFirewall_updateFirewallDeleteProtectionCmd.Flags().String("update-token", "", "An optional token that you can use for optimistic locking.")
		networkFirewall_updateFirewallDeleteProtectionCmd.MarkFlagRequired("delete-protection")
		networkFirewall_updateFirewallDeleteProtectionCmd.Flag("no-delete-protection").Hidden = true
		networkFirewall_updateFirewallDeleteProtectionCmd.MarkFlagRequired("no-delete-protection")
	})
	networkFirewallCmd.AddCommand(networkFirewall_updateFirewallDeleteProtectionCmd)
}
