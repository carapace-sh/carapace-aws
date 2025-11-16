package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var networkFirewall_disassociateAvailabilityZonesCmd = &cobra.Command{
	Use:   "disassociate-availability-zones",
	Short: "Removes the specified Availability Zone associations from a transit gateway-attached firewall.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(networkFirewall_disassociateAvailabilityZonesCmd).Standalone()

	networkFirewall_disassociateAvailabilityZonesCmd.Flags().String("availability-zone-mappings", "", "Required.")
	networkFirewall_disassociateAvailabilityZonesCmd.Flags().String("firewall-arn", "", "The Amazon Resource Name (ARN) of the firewall.")
	networkFirewall_disassociateAvailabilityZonesCmd.Flags().String("firewall-name", "", "The descriptive name of the firewall.")
	networkFirewall_disassociateAvailabilityZonesCmd.Flags().String("update-token", "", "An optional token that you can use for optimistic locking.")
	networkFirewall_disassociateAvailabilityZonesCmd.MarkFlagRequired("availability-zone-mappings")
	networkFirewallCmd.AddCommand(networkFirewall_disassociateAvailabilityZonesCmd)
}
