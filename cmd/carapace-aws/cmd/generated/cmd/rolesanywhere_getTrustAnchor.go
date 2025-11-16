package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rolesanywhere_getTrustAnchorCmd = &cobra.Command{
	Use:   "get-trust-anchor",
	Short: "Gets a trust anchor.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rolesanywhere_getTrustAnchorCmd).Standalone()

	rolesanywhere_getTrustAnchorCmd.Flags().String("trust-anchor-id", "", "The unique identifier of the trust anchor.")
	rolesanywhere_getTrustAnchorCmd.MarkFlagRequired("trust-anchor-id")
	rolesanywhereCmd.AddCommand(rolesanywhere_getTrustAnchorCmd)
}
