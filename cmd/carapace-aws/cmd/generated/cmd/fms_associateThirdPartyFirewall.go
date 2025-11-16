package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var fms_associateThirdPartyFirewallCmd = &cobra.Command{
	Use:   "associate-third-party-firewall",
	Short: "Sets the Firewall Manager policy administrator as a tenant administrator of a third-party firewall service.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(fms_associateThirdPartyFirewallCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(fms_associateThirdPartyFirewallCmd).Standalone()

		fms_associateThirdPartyFirewallCmd.Flags().String("third-party-firewall", "", "The name of the third-party firewall vendor.")
		fms_associateThirdPartyFirewallCmd.MarkFlagRequired("third-party-firewall")
	})
	fmsCmd.AddCommand(fms_associateThirdPartyFirewallCmd)
}
