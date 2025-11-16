package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var networkFirewall_describeLoggingConfigurationCmd = &cobra.Command{
	Use:   "describe-logging-configuration",
	Short: "Returns the logging configuration for the specified firewall.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(networkFirewall_describeLoggingConfigurationCmd).Standalone()

	networkFirewall_describeLoggingConfigurationCmd.Flags().String("firewall-arn", "", "The Amazon Resource Name (ARN) of the firewall.")
	networkFirewall_describeLoggingConfigurationCmd.Flags().String("firewall-name", "", "The descriptive name of the firewall.")
	networkFirewallCmd.AddCommand(networkFirewall_describeLoggingConfigurationCmd)
}
