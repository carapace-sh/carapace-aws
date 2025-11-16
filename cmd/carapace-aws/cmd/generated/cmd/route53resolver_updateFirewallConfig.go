package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53resolver_updateFirewallConfigCmd = &cobra.Command{
	Use:   "update-firewall-config",
	Short: "Updates the configuration of the firewall behavior provided by DNS Firewall for a single VPC from Amazon Virtual Private Cloud (Amazon VPC).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53resolver_updateFirewallConfigCmd).Standalone()

	route53resolver_updateFirewallConfigCmd.Flags().String("firewall-fail-open", "", "Determines how Route 53 Resolver handles queries during failures, for example when all traffic that is sent to DNS Firewall fails to receive a reply.")
	route53resolver_updateFirewallConfigCmd.Flags().String("resource-id", "", "The ID of the VPC that the configuration is for.")
	route53resolver_updateFirewallConfigCmd.MarkFlagRequired("firewall-fail-open")
	route53resolver_updateFirewallConfigCmd.MarkFlagRequired("resource-id")
	route53resolverCmd.AddCommand(route53resolver_updateFirewallConfigCmd)
}
