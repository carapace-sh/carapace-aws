package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var fms_disassociateThirdPartyFirewallCmd = &cobra.Command{
	Use:   "disassociate-third-party-firewall",
	Short: "Disassociates a Firewall Manager policy administrator from a third-party firewall tenant.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(fms_disassociateThirdPartyFirewallCmd).Standalone()

	fms_disassociateThirdPartyFirewallCmd.Flags().String("third-party-firewall", "", "The name of the third-party firewall vendor.")
	fms_disassociateThirdPartyFirewallCmd.MarkFlagRequired("third-party-firewall")
	fmsCmd.AddCommand(fms_disassociateThirdPartyFirewallCmd)
}
