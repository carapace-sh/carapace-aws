package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var networkFirewall_createFirewallCmd = &cobra.Command{
	Use:   "create-firewall",
	Short: "Creates an Network Firewall [Firewall]() and accompanying [FirewallStatus]() for a VPC.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(networkFirewall_createFirewallCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(networkFirewall_createFirewallCmd).Standalone()

		networkFirewall_createFirewallCmd.Flags().Bool("availability-zone-change-protection", false, "Optional.")
		networkFirewall_createFirewallCmd.Flags().String("availability-zone-mappings", "", "Required.")
		networkFirewall_createFirewallCmd.Flags().Bool("delete-protection", false, "A flag indicating whether it is possible to delete the firewall.")
		networkFirewall_createFirewallCmd.Flags().String("description", "", "A description of the firewall.")
		networkFirewall_createFirewallCmd.Flags().String("enabled-analysis-types", "", "An optional setting indicating the specific traffic analysis types to enable on the firewall.")
		networkFirewall_createFirewallCmd.Flags().String("encryption-configuration", "", "A complex type that contains settings for encryption of your firewall resources.")
		networkFirewall_createFirewallCmd.Flags().String("firewall-name", "", "The descriptive name of the firewall.")
		networkFirewall_createFirewallCmd.Flags().String("firewall-policy-arn", "", "The Amazon Resource Name (ARN) of the [FirewallPolicy]() that you want to use for the firewall.")
		networkFirewall_createFirewallCmd.Flags().Bool("firewall-policy-change-protection", false, "A setting indicating whether the firewall is protected against a change to the firewall policy association.")
		networkFirewall_createFirewallCmd.Flags().Bool("no-availability-zone-change-protection", false, "Optional.")
		networkFirewall_createFirewallCmd.Flags().Bool("no-delete-protection", false, "A flag indicating whether it is possible to delete the firewall.")
		networkFirewall_createFirewallCmd.Flags().Bool("no-firewall-policy-change-protection", false, "A setting indicating whether the firewall is protected against a change to the firewall policy association.")
		networkFirewall_createFirewallCmd.Flags().Bool("no-subnet-change-protection", false, "A setting indicating whether the firewall is protected against changes to the subnet associations.")
		networkFirewall_createFirewallCmd.Flags().Bool("subnet-change-protection", false, "A setting indicating whether the firewall is protected against changes to the subnet associations.")
		networkFirewall_createFirewallCmd.Flags().String("subnet-mappings", "", "The public subnets to use for your Network Firewall firewalls.")
		networkFirewall_createFirewallCmd.Flags().String("tags", "", "The key:value pairs to associate with the resource.")
		networkFirewall_createFirewallCmd.Flags().String("transit-gateway-id", "", "Required when creating a transit gateway-attached firewall.")
		networkFirewall_createFirewallCmd.Flags().String("vpc-id", "", "The unique identifier of the VPC where Network Firewall should create the firewall.")
		networkFirewall_createFirewallCmd.MarkFlagRequired("firewall-name")
		networkFirewall_createFirewallCmd.MarkFlagRequired("firewall-policy-arn")
		networkFirewall_createFirewallCmd.Flag("no-availability-zone-change-protection").Hidden = true
		networkFirewall_createFirewallCmd.Flag("no-delete-protection").Hidden = true
		networkFirewall_createFirewallCmd.Flag("no-firewall-policy-change-protection").Hidden = true
		networkFirewall_createFirewallCmd.Flag("no-subnet-change-protection").Hidden = true
	})
	networkFirewallCmd.AddCommand(networkFirewall_createFirewallCmd)
}
