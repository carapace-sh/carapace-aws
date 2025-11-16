package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rolesanywhere_enableTrustAnchorCmd = &cobra.Command{
	Use:   "enable-trust-anchor",
	Short: "Enables a trust anchor.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rolesanywhere_enableTrustAnchorCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(rolesanywhere_enableTrustAnchorCmd).Standalone()

		rolesanywhere_enableTrustAnchorCmd.Flags().String("trust-anchor-id", "", "The unique identifier of the trust anchor.")
		rolesanywhere_enableTrustAnchorCmd.MarkFlagRequired("trust-anchor-id")
	})
	rolesanywhereCmd.AddCommand(rolesanywhere_enableTrustAnchorCmd)
}
