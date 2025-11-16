package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var networkFirewall_updateFirewallPolicyCmd = &cobra.Command{
	Use:   "update-firewall-policy",
	Short: "Updates the properties of the specified firewall policy.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(networkFirewall_updateFirewallPolicyCmd).Standalone()

	networkFirewall_updateFirewallPolicyCmd.Flags().String("description", "", "A description of the firewall policy.")
	networkFirewall_updateFirewallPolicyCmd.Flags().Bool("dry-run", false, "Indicates whether you want Network Firewall to just check the validity of the request, rather than run the request.")
	networkFirewall_updateFirewallPolicyCmd.Flags().String("encryption-configuration", "", "A complex type that contains settings for encryption of your firewall policy resources.")
	networkFirewall_updateFirewallPolicyCmd.Flags().String("firewall-policy", "", "The updated firewall policy to use for the firewall.")
	networkFirewall_updateFirewallPolicyCmd.Flags().String("firewall-policy-arn", "", "The Amazon Resource Name (ARN) of the firewall policy.")
	networkFirewall_updateFirewallPolicyCmd.Flags().String("firewall-policy-name", "", "The descriptive name of the firewall policy.")
	networkFirewall_updateFirewallPolicyCmd.Flags().Bool("no-dry-run", false, "Indicates whether you want Network Firewall to just check the validity of the request, rather than run the request.")
	networkFirewall_updateFirewallPolicyCmd.Flags().String("update-token", "", "A token used for optimistic locking.")
	networkFirewall_updateFirewallPolicyCmd.MarkFlagRequired("firewall-policy")
	networkFirewall_updateFirewallPolicyCmd.Flag("no-dry-run").Hidden = true
	networkFirewall_updateFirewallPolicyCmd.MarkFlagRequired("update-token")
	networkFirewallCmd.AddCommand(networkFirewall_updateFirewallPolicyCmd)
}
