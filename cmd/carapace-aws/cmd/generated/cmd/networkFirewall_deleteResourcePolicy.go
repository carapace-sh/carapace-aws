package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var networkFirewall_deleteResourcePolicyCmd = &cobra.Command{
	Use:   "delete-resource-policy",
	Short: "Deletes a resource policy that you created in a [PutResourcePolicy]() request.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(networkFirewall_deleteResourcePolicyCmd).Standalone()

	networkFirewall_deleteResourcePolicyCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the rule group or firewall policy whose resource policy you want to delete.")
	networkFirewall_deleteResourcePolicyCmd.MarkFlagRequired("resource-arn")
	networkFirewallCmd.AddCommand(networkFirewall_deleteResourcePolicyCmd)
}
