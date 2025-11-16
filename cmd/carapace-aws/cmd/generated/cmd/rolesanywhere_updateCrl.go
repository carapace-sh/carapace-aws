package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rolesanywhere_updateCrlCmd = &cobra.Command{
	Use:   "update-crl",
	Short: "Updates the certificate revocation list (CRL).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rolesanywhere_updateCrlCmd).Standalone()

	rolesanywhere_updateCrlCmd.Flags().String("crl-data", "", "The x509 v3 specified certificate revocation list (CRL).")
	rolesanywhere_updateCrlCmd.Flags().String("crl-id", "", "The unique identifier of the certificate revocation list (CRL).")
	rolesanywhere_updateCrlCmd.Flags().String("name", "", "The name of the Crl.")
	rolesanywhere_updateCrlCmd.MarkFlagRequired("crl-id")
	rolesanywhereCmd.AddCommand(rolesanywhere_updateCrlCmd)
}
