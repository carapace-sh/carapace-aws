package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var networkFirewall_deleteTlsinspectionConfigurationCmd = &cobra.Command{
	Use:   "delete-tlsinspection-configuration",
	Short: "Deletes the specified [TLSInspectionConfiguration]().",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(networkFirewall_deleteTlsinspectionConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(networkFirewall_deleteTlsinspectionConfigurationCmd).Standalone()

		networkFirewall_deleteTlsinspectionConfigurationCmd.Flags().String("tlsinspection-configuration-arn", "", "The Amazon Resource Name (ARN) of the TLS inspection configuration.")
		networkFirewall_deleteTlsinspectionConfigurationCmd.Flags().String("tlsinspection-configuration-name", "", "The descriptive name of the TLS inspection configuration.")
	})
	networkFirewallCmd.AddCommand(networkFirewall_deleteTlsinspectionConfigurationCmd)
}
