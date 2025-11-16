package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var networkFirewall_updateFirewallAnalysisSettingsCmd = &cobra.Command{
	Use:   "update-firewall-analysis-settings",
	Short: "Enables specific types of firewall analysis on a specific firewall you define.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(networkFirewall_updateFirewallAnalysisSettingsCmd).Standalone()

	networkFirewall_updateFirewallAnalysisSettingsCmd.Flags().String("enabled-analysis-types", "", "An optional setting indicating the specific traffic analysis types to enable on the firewall.")
	networkFirewall_updateFirewallAnalysisSettingsCmd.Flags().String("firewall-arn", "", "The Amazon Resource Name (ARN) of the firewall.")
	networkFirewall_updateFirewallAnalysisSettingsCmd.Flags().String("firewall-name", "", "The descriptive name of the firewall.")
	networkFirewall_updateFirewallAnalysisSettingsCmd.Flags().String("update-token", "", "An optional token that you can use for optimistic locking.")
	networkFirewallCmd.AddCommand(networkFirewall_updateFirewallAnalysisSettingsCmd)
}
