package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var fms_getThirdPartyFirewallAssociationStatusCmd = &cobra.Command{
	Use:   "get-third-party-firewall-association-status",
	Short: "The onboarding status of a Firewall Manager admin account to third-party firewall vendor tenant.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(fms_getThirdPartyFirewallAssociationStatusCmd).Standalone()

	fms_getThirdPartyFirewallAssociationStatusCmd.Flags().String("third-party-firewall", "", "The name of the third-party firewall vendor.")
	fms_getThirdPartyFirewallAssociationStatusCmd.MarkFlagRequired("third-party-firewall")
	fmsCmd.AddCommand(fms_getThirdPartyFirewallAssociationStatusCmd)
}
