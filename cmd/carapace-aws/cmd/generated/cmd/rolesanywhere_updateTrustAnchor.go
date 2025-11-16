package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rolesanywhere_updateTrustAnchorCmd = &cobra.Command{
	Use:   "update-trust-anchor",
	Short: "Updates a trust anchor.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rolesanywhere_updateTrustAnchorCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(rolesanywhere_updateTrustAnchorCmd).Standalone()

		rolesanywhere_updateTrustAnchorCmd.Flags().String("name", "", "The name of the trust anchor.")
		rolesanywhere_updateTrustAnchorCmd.Flags().String("source", "", "The trust anchor type and its related certificate data.")
		rolesanywhere_updateTrustAnchorCmd.Flags().String("trust-anchor-id", "", "The unique identifier of the trust anchor.")
		rolesanywhere_updateTrustAnchorCmd.MarkFlagRequired("trust-anchor-id")
	})
	rolesanywhereCmd.AddCommand(rolesanywhere_updateTrustAnchorCmd)
}
