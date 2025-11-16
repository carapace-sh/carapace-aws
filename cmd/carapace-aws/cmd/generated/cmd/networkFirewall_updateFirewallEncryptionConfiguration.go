package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var networkFirewall_updateFirewallEncryptionConfigurationCmd = &cobra.Command{
	Use:   "update-firewall-encryption-configuration",
	Short: "A complex type that contains settings for encryption of your firewall resources.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(networkFirewall_updateFirewallEncryptionConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(networkFirewall_updateFirewallEncryptionConfigurationCmd).Standalone()

		networkFirewall_updateFirewallEncryptionConfigurationCmd.Flags().String("encryption-configuration", "", "")
		networkFirewall_updateFirewallEncryptionConfigurationCmd.Flags().String("firewall-arn", "", "The Amazon Resource Name (ARN) of the firewall.")
		networkFirewall_updateFirewallEncryptionConfigurationCmd.Flags().String("firewall-name", "", "The descriptive name of the firewall.")
		networkFirewall_updateFirewallEncryptionConfigurationCmd.Flags().String("update-token", "", "An optional token that you can use for optimistic locking.")
	})
	networkFirewallCmd.AddCommand(networkFirewall_updateFirewallEncryptionConfigurationCmd)
}
