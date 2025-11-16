package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var networkFirewall_disassociateSubnetsCmd = &cobra.Command{
	Use:   "disassociate-subnets",
	Short: "Removes the specified subnet associations from the firewall.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(networkFirewall_disassociateSubnetsCmd).Standalone()

	networkFirewall_disassociateSubnetsCmd.Flags().String("firewall-arn", "", "The Amazon Resource Name (ARN) of the firewall.")
	networkFirewall_disassociateSubnetsCmd.Flags().String("firewall-name", "", "The descriptive name of the firewall.")
	networkFirewall_disassociateSubnetsCmd.Flags().String("subnet-ids", "", "The unique identifiers for the subnets that you want to disassociate.")
	networkFirewall_disassociateSubnetsCmd.Flags().String("update-token", "", "An optional token that you can use for optimistic locking.")
	networkFirewall_disassociateSubnetsCmd.MarkFlagRequired("subnet-ids")
	networkFirewallCmd.AddCommand(networkFirewall_disassociateSubnetsCmd)
}
