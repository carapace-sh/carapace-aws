package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var networkFirewall_updateAvailabilityZoneChangeProtectionCmd = &cobra.Command{
	Use:   "update-availability-zone-change-protection",
	Short: "Modifies the `AvailabilityZoneChangeProtection` setting for a transit gateway-attached firewall.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(networkFirewall_updateAvailabilityZoneChangeProtectionCmd).Standalone()

	networkFirewall_updateAvailabilityZoneChangeProtectionCmd.Flags().Bool("availability-zone-change-protection", false, "A setting indicating whether the firewall is protected against changes to the subnet associations.")
	networkFirewall_updateAvailabilityZoneChangeProtectionCmd.Flags().String("firewall-arn", "", "The Amazon Resource Name (ARN) of the firewall.")
	networkFirewall_updateAvailabilityZoneChangeProtectionCmd.Flags().String("firewall-name", "", "The descriptive name of the firewall.")
	networkFirewall_updateAvailabilityZoneChangeProtectionCmd.Flags().Bool("no-availability-zone-change-protection", false, "A setting indicating whether the firewall is protected against changes to the subnet associations.")
	networkFirewall_updateAvailabilityZoneChangeProtectionCmd.Flags().String("update-token", "", "An optional token that you can use for optimistic locking.")
	networkFirewall_updateAvailabilityZoneChangeProtectionCmd.MarkFlagRequired("availability-zone-change-protection")
	networkFirewall_updateAvailabilityZoneChangeProtectionCmd.Flag("no-availability-zone-change-protection").Hidden = true
	networkFirewall_updateAvailabilityZoneChangeProtectionCmd.MarkFlagRequired("no-availability-zone-change-protection")
	networkFirewallCmd.AddCommand(networkFirewall_updateAvailabilityZoneChangeProtectionCmd)
}
