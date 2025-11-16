package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var networkFirewall_describeFirewallMetadataCmd = &cobra.Command{
	Use:   "describe-firewall-metadata",
	Short: "Returns the high-level information about a firewall, including the Availability Zones where the Firewall is currently in use.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(networkFirewall_describeFirewallMetadataCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(networkFirewall_describeFirewallMetadataCmd).Standalone()

		networkFirewall_describeFirewallMetadataCmd.Flags().String("firewall-arn", "", "The Amazon Resource Name (ARN) of the firewall.")
	})
	networkFirewallCmd.AddCommand(networkFirewall_describeFirewallMetadataCmd)
}
