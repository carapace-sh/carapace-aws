package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rolesanywhere_enableCrlCmd = &cobra.Command{
	Use:   "enable-crl",
	Short: "Enables a certificate revocation list (CRL).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rolesanywhere_enableCrlCmd).Standalone()

	rolesanywhere_enableCrlCmd.Flags().String("crl-id", "", "The unique identifier of the certificate revocation list (CRL).")
	rolesanywhere_enableCrlCmd.MarkFlagRequired("crl-id")
	rolesanywhereCmd.AddCommand(rolesanywhere_enableCrlCmd)
}
