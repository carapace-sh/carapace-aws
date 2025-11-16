package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rolesanywhere_deleteTrustAnchorCmd = &cobra.Command{
	Use:   "delete-trust-anchor",
	Short: "Deletes a trust anchor.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rolesanywhere_deleteTrustAnchorCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(rolesanywhere_deleteTrustAnchorCmd).Standalone()

		rolesanywhere_deleteTrustAnchorCmd.Flags().String("trust-anchor-id", "", "The unique identifier of the trust anchor.")
		rolesanywhere_deleteTrustAnchorCmd.MarkFlagRequired("trust-anchor-id")
	})
	rolesanywhereCmd.AddCommand(rolesanywhere_deleteTrustAnchorCmd)
}
