package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var acmPca_updateCertificateAuthorityCmd = &cobra.Command{
	Use:   "update-certificate-authority",
	Short: "Updates the status or configuration of a private certificate authority (CA).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(acmPca_updateCertificateAuthorityCmd).Standalone()

	acmPca_updateCertificateAuthorityCmd.Flags().String("certificate-authority-arn", "", "Amazon Resource Name (ARN) of the private CA that issued the certificate to be revoked.")
	acmPca_updateCertificateAuthorityCmd.Flags().String("revocation-configuration", "", "Contains information to enable support for Online Certificate Status Protocol (OCSP), certificate revocation list (CRL), both protocols, or neither.")
	acmPca_updateCertificateAuthorityCmd.Flags().String("status", "", "Status of your private CA.")
	acmPca_updateCertificateAuthorityCmd.MarkFlagRequired("certificate-authority-arn")
	acmPcaCmd.AddCommand(acmPca_updateCertificateAuthorityCmd)
}
