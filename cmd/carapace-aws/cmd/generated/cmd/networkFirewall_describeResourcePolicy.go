package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var networkFirewall_describeResourcePolicyCmd = &cobra.Command{
	Use:   "describe-resource-policy",
	Short: "Retrieves a resource policy that you created in a [PutResourcePolicy]() request.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(networkFirewall_describeResourcePolicyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(networkFirewall_describeResourcePolicyCmd).Standalone()

		networkFirewall_describeResourcePolicyCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the rule group or firewall policy whose resource policy you want to retrieve.")
		networkFirewall_describeResourcePolicyCmd.MarkFlagRequired("resource-arn")
	})
	networkFirewallCmd.AddCommand(networkFirewall_describeResourcePolicyCmd)
}
