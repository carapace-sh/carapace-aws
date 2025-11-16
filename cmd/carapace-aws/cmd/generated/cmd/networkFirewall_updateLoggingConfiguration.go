package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var networkFirewall_updateLoggingConfigurationCmd = &cobra.Command{
	Use:   "update-logging-configuration",
	Short: "Sets the logging configuration for the specified firewall.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(networkFirewall_updateLoggingConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(networkFirewall_updateLoggingConfigurationCmd).Standalone()

		networkFirewall_updateLoggingConfigurationCmd.Flags().String("enable-monitoring-dashboard", "", "A boolean that lets you enable or disable the detailed firewall monitoring dashboard on the firewall.")
		networkFirewall_updateLoggingConfigurationCmd.Flags().String("firewall-arn", "", "The Amazon Resource Name (ARN) of the firewall.")
		networkFirewall_updateLoggingConfigurationCmd.Flags().String("firewall-name", "", "The descriptive name of the firewall.")
		networkFirewall_updateLoggingConfigurationCmd.Flags().String("logging-configuration", "", "Defines how Network Firewall performs logging for a firewall.")
	})
	networkFirewallCmd.AddCommand(networkFirewall_updateLoggingConfigurationCmd)
}
