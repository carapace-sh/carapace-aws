package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var networkFirewall_describeTlsinspectionConfigurationCmd = &cobra.Command{
	Use:   "describe-tlsinspection-configuration",
	Short: "Returns the data objects for the specified TLS inspection configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(networkFirewall_describeTlsinspectionConfigurationCmd).Standalone()

	networkFirewall_describeTlsinspectionConfigurationCmd.Flags().String("tlsinspection-configuration-arn", "", "The Amazon Resource Name (ARN) of the TLS inspection configuration.")
	networkFirewall_describeTlsinspectionConfigurationCmd.Flags().String("tlsinspection-configuration-name", "", "The descriptive name of the TLS inspection configuration.")
	networkFirewallCmd.AddCommand(networkFirewall_describeTlsinspectionConfigurationCmd)
}
