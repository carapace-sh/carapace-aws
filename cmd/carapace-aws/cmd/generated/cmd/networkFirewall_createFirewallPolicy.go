package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var networkFirewall_createFirewallPolicyCmd = &cobra.Command{
	Use:   "create-firewall-policy",
	Short: "Creates the firewall policy for the firewall according to the specifications.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(networkFirewall_createFirewallPolicyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(networkFirewall_createFirewallPolicyCmd).Standalone()

		networkFirewall_createFirewallPolicyCmd.Flags().String("description", "", "A description of the firewall policy.")
		networkFirewall_createFirewallPolicyCmd.Flags().Bool("dry-run", false, "Indicates whether you want Network Firewall to just check the validity of the request, rather than run the request.")
		networkFirewall_createFirewallPolicyCmd.Flags().String("encryption-configuration", "", "A complex type that contains settings for encryption of your firewall policy resources.")
		networkFirewall_createFirewallPolicyCmd.Flags().String("firewall-policy", "", "The rule groups and policy actions to use in the firewall policy.")
		networkFirewall_createFirewallPolicyCmd.Flags().String("firewall-policy-name", "", "The descriptive name of the firewall policy.")
		networkFirewall_createFirewallPolicyCmd.Flags().Bool("no-dry-run", false, "Indicates whether you want Network Firewall to just check the validity of the request, rather than run the request.")
		networkFirewall_createFirewallPolicyCmd.Flags().String("tags", "", "The key:value pairs to associate with the resource.")
		networkFirewall_createFirewallPolicyCmd.MarkFlagRequired("firewall-policy")
		networkFirewall_createFirewallPolicyCmd.MarkFlagRequired("firewall-policy-name")
		networkFirewall_createFirewallPolicyCmd.Flag("no-dry-run").Hidden = true
	})
	networkFirewallCmd.AddCommand(networkFirewall_createFirewallPolicyCmd)
}
