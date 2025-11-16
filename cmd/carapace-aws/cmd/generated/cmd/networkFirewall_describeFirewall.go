package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var networkFirewall_describeFirewallCmd = &cobra.Command{
	Use:   "describe-firewall",
	Short: "Returns the data objects for the specified firewall.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(networkFirewall_describeFirewallCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(networkFirewall_describeFirewallCmd).Standalone()

		networkFirewall_describeFirewallCmd.Flags().String("firewall-arn", "", "The Amazon Resource Name (ARN) of the firewall.")
		networkFirewall_describeFirewallCmd.Flags().String("firewall-name", "", "The descriptive name of the firewall.")
	})
	networkFirewallCmd.AddCommand(networkFirewall_describeFirewallCmd)
}
