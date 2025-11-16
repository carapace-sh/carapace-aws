package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var networkFirewall_updateFirewallDescriptionCmd = &cobra.Command{
	Use:   "update-firewall-description",
	Short: "Modifies the description for the specified firewall.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(networkFirewall_updateFirewallDescriptionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(networkFirewall_updateFirewallDescriptionCmd).Standalone()

		networkFirewall_updateFirewallDescriptionCmd.Flags().String("description", "", "The new description for the firewall.")
		networkFirewall_updateFirewallDescriptionCmd.Flags().String("firewall-arn", "", "The Amazon Resource Name (ARN) of the firewall.")
		networkFirewall_updateFirewallDescriptionCmd.Flags().String("firewall-name", "", "The descriptive name of the firewall.")
		networkFirewall_updateFirewallDescriptionCmd.Flags().String("update-token", "", "An optional token that you can use for optimistic locking.")
	})
	networkFirewallCmd.AddCommand(networkFirewall_updateFirewallDescriptionCmd)
}
