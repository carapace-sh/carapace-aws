package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rolesanywhere_deleteCrlCmd = &cobra.Command{
	Use:   "delete-crl",
	Short: "Deletes a certificate revocation list (CRL).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rolesanywhere_deleteCrlCmd).Standalone()

	rolesanywhere_deleteCrlCmd.Flags().String("crl-id", "", "The unique identifier of the certificate revocation list (CRL).")
	rolesanywhere_deleteCrlCmd.MarkFlagRequired("crl-id")
	rolesanywhereCmd.AddCommand(rolesanywhere_deleteCrlCmd)
}
