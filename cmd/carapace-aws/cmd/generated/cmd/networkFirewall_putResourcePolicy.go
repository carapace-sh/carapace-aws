package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var networkFirewall_putResourcePolicyCmd = &cobra.Command{
	Use:   "put-resource-policy",
	Short: "Creates or updates an IAM policy for your rule group, firewall policy, or firewall.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(networkFirewall_putResourcePolicyCmd).Standalone()

	networkFirewall_putResourcePolicyCmd.Flags().String("policy", "", "The IAM policy statement that lists the accounts that you want to share your Network Firewall resources with and the operations that you want the accounts to be able to perform.")
	networkFirewall_putResourcePolicyCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the account that you want to share your Network Firewall resources with.")
	networkFirewall_putResourcePolicyCmd.MarkFlagRequired("policy")
	networkFirewall_putResourcePolicyCmd.MarkFlagRequired("resource-arn")
	networkFirewallCmd.AddCommand(networkFirewall_putResourcePolicyCmd)
}
