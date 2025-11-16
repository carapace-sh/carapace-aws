package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rolesanywhere_disableTrustAnchorCmd = &cobra.Command{
	Use:   "disable-trust-anchor",
	Short: "Disables a trust anchor.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rolesanywhere_disableTrustAnchorCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(rolesanywhere_disableTrustAnchorCmd).Standalone()

		rolesanywhere_disableTrustAnchorCmd.Flags().String("trust-anchor-id", "", "The unique identifier of the trust anchor.")
		rolesanywhere_disableTrustAnchorCmd.MarkFlagRequired("trust-anchor-id")
	})
	rolesanywhereCmd.AddCommand(rolesanywhere_disableTrustAnchorCmd)
}
