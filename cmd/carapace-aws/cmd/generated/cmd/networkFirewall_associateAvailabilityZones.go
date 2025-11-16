package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var networkFirewall_associateAvailabilityZonesCmd = &cobra.Command{
	Use:   "associate-availability-zones",
	Short: "Associates the specified Availability Zones with a transit gateway-attached firewall.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(networkFirewall_associateAvailabilityZonesCmd).Standalone()

	networkFirewall_associateAvailabilityZonesCmd.Flags().String("availability-zone-mappings", "", "Required.")
	networkFirewall_associateAvailabilityZonesCmd.Flags().String("firewall-arn", "", "The Amazon Resource Name (ARN) of the firewall.")
	networkFirewall_associateAvailabilityZonesCmd.Flags().String("firewall-name", "", "The descriptive name of the firewall.")
	networkFirewall_associateAvailabilityZonesCmd.Flags().String("update-token", "", "An optional token that you can use for optimistic locking.")
	networkFirewall_associateAvailabilityZonesCmd.MarkFlagRequired("availability-zone-mappings")
	networkFirewallCmd.AddCommand(networkFirewall_associateAvailabilityZonesCmd)
}
