package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var networkFirewall_deleteFirewallCmd = &cobra.Command{
	Use:   "delete-firewall",
	Short: "Deletes the specified [Firewall]() and its [FirewallStatus]().",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(networkFirewall_deleteFirewallCmd).Standalone()

	networkFirewall_deleteFirewallCmd.Flags().String("firewall-arn", "", "The Amazon Resource Name (ARN) of the firewall.")
	networkFirewall_deleteFirewallCmd.Flags().String("firewall-name", "", "The descriptive name of the firewall.")
	networkFirewallCmd.AddCommand(networkFirewall_deleteFirewallCmd)
}
