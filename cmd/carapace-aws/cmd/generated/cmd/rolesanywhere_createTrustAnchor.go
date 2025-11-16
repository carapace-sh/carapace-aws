package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rolesanywhere_createTrustAnchorCmd = &cobra.Command{
	Use:   "create-trust-anchor",
	Short: "Creates a trust anchor to establish trust between IAM Roles Anywhere and your certificate authority (CA).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rolesanywhere_createTrustAnchorCmd).Standalone()

	rolesanywhere_createTrustAnchorCmd.Flags().Bool("enabled", false, "Specifies whether the trust anchor is enabled.")
	rolesanywhere_createTrustAnchorCmd.Flags().String("name", "", "The name of the trust anchor.")
	rolesanywhere_createTrustAnchorCmd.Flags().Bool("no-enabled", false, "Specifies whether the trust anchor is enabled.")
	rolesanywhere_createTrustAnchorCmd.Flags().String("notification-settings", "", "A list of notification settings to be associated to the trust anchor.")
	rolesanywhere_createTrustAnchorCmd.Flags().String("source", "", "The trust anchor type and its related certificate data.")
	rolesanywhere_createTrustAnchorCmd.Flags().String("tags", "", "The tags to attach to the trust anchor.")
	rolesanywhere_createTrustAnchorCmd.MarkFlagRequired("name")
	rolesanywhere_createTrustAnchorCmd.Flag("no-enabled").Hidden = true
	rolesanywhere_createTrustAnchorCmd.MarkFlagRequired("source")
	rolesanywhereCmd.AddCommand(rolesanywhere_createTrustAnchorCmd)
}
