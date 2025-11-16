package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53resolver_getFirewallConfigCmd = &cobra.Command{
	Use:   "get-firewall-config",
	Short: "Retrieves the configuration of the firewall behavior provided by DNS Firewall for a single VPC from Amazon Virtual Private Cloud (Amazon VPC).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53resolver_getFirewallConfigCmd).Standalone()

	route53resolver_getFirewallConfigCmd.Flags().String("resource-id", "", "The ID of the VPC from Amazon VPC that the configuration is for.")
	route53resolver_getFirewallConfigCmd.MarkFlagRequired("resource-id")
	route53resolverCmd.AddCommand(route53resolver_getFirewallConfigCmd)
}
