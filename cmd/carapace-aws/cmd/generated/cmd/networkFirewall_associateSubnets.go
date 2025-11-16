package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var networkFirewall_associateSubnetsCmd = &cobra.Command{
	Use:   "associate-subnets",
	Short: "Associates the specified subnets in the Amazon VPC to the firewall.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(networkFirewall_associateSubnetsCmd).Standalone()

	networkFirewall_associateSubnetsCmd.Flags().String("firewall-arn", "", "The Amazon Resource Name (ARN) of the firewall.")
	networkFirewall_associateSubnetsCmd.Flags().String("firewall-name", "", "The descriptive name of the firewall.")
	networkFirewall_associateSubnetsCmd.Flags().String("subnet-mappings", "", "The IDs of the subnets that you want to associate with the firewall.")
	networkFirewall_associateSubnetsCmd.Flags().String("update-token", "", "An optional token that you can use for optimistic locking.")
	networkFirewall_associateSubnetsCmd.MarkFlagRequired("subnet-mappings")
	networkFirewallCmd.AddCommand(networkFirewall_associateSubnetsCmd)
}
