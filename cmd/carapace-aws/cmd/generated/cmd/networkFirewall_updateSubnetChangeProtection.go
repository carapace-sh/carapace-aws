package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var networkFirewall_updateSubnetChangeProtectionCmd = &cobra.Command{
	Use:   "update-subnet-change-protection",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(networkFirewall_updateSubnetChangeProtectionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(networkFirewall_updateSubnetChangeProtectionCmd).Standalone()

		networkFirewall_updateSubnetChangeProtectionCmd.Flags().String("firewall-arn", "", "The Amazon Resource Name (ARN) of the firewall.")
		networkFirewall_updateSubnetChangeProtectionCmd.Flags().String("firewall-name", "", "The descriptive name of the firewall.")
		networkFirewall_updateSubnetChangeProtectionCmd.Flags().Bool("no-subnet-change-protection", false, "A setting indicating whether the firewall is protected against changes to the subnet associations.")
		networkFirewall_updateSubnetChangeProtectionCmd.Flags().Bool("subnet-change-protection", false, "A setting indicating whether the firewall is protected against changes to the subnet associations.")
		networkFirewall_updateSubnetChangeProtectionCmd.Flags().String("update-token", "", "An optional token that you can use for optimistic locking.")
		networkFirewall_updateSubnetChangeProtectionCmd.Flag("no-subnet-change-protection").Hidden = true
		networkFirewall_updateSubnetChangeProtectionCmd.MarkFlagRequired("no-subnet-change-protection")
		networkFirewall_updateSubnetChangeProtectionCmd.MarkFlagRequired("subnet-change-protection")
	})
	networkFirewallCmd.AddCommand(networkFirewall_updateSubnetChangeProtectionCmd)
}
