package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var networkFirewall_createTlsinspectionConfigurationCmd = &cobra.Command{
	Use:   "create-tlsinspection-configuration",
	Short: "Creates an Network Firewall TLS inspection configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(networkFirewall_createTlsinspectionConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(networkFirewall_createTlsinspectionConfigurationCmd).Standalone()

		networkFirewall_createTlsinspectionConfigurationCmd.Flags().String("description", "", "A description of the TLS inspection configuration.")
		networkFirewall_createTlsinspectionConfigurationCmd.Flags().String("encryption-configuration", "", "")
		networkFirewall_createTlsinspectionConfigurationCmd.Flags().String("tags", "", "The key:value pairs to associate with the resource.")
		networkFirewall_createTlsinspectionConfigurationCmd.Flags().String("tlsinspection-configuration", "", "The object that defines a TLS inspection configuration.")
		networkFirewall_createTlsinspectionConfigurationCmd.Flags().String("tlsinspection-configuration-name", "", "The descriptive name of the TLS inspection configuration.")
		networkFirewall_createTlsinspectionConfigurationCmd.MarkFlagRequired("tlsinspection-configuration")
		networkFirewall_createTlsinspectionConfigurationCmd.MarkFlagRequired("tlsinspection-configuration-name")
	})
	networkFirewallCmd.AddCommand(networkFirewall_createTlsinspectionConfigurationCmd)
}
