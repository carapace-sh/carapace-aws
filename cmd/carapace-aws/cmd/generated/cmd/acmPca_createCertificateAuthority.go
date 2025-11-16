package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var acmPca_createCertificateAuthorityCmd = &cobra.Command{
	Use:   "create-certificate-authority",
	Short: "Creates a root or subordinate private certificate authority (CA).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(acmPca_createCertificateAuthorityCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(acmPca_createCertificateAuthorityCmd).Standalone()

		acmPca_createCertificateAuthorityCmd.Flags().String("certificate-authority-configuration", "", "Name and bit size of the private key algorithm, the name of the signing algorithm, and X.500 certificate subject information.")
		acmPca_createCertificateAuthorityCmd.Flags().String("certificate-authority-type", "", "The type of the certificate authority.")
		acmPca_createCertificateAuthorityCmd.Flags().String("idempotency-token", "", "Custom string that can be used to distinguish between calls to the **CreateCertificateAuthority** action.")
		acmPca_createCertificateAuthorityCmd.Flags().String("key-storage-security-standard", "", "Specifies a cryptographic key management compliance standard for handling and protecting CA keys.")
		acmPca_createCertificateAuthorityCmd.Flags().String("revocation-configuration", "", "Contains information to enable support for Online Certificate Status Protocol (OCSP), certificate revocation list (CRL), both protocols, or neither.")
		acmPca_createCertificateAuthorityCmd.Flags().String("tags", "", "Key-value pairs that will be attached to the new private CA.")
		acmPca_createCertificateAuthorityCmd.Flags().String("usage-mode", "", "Specifies whether the CA issues general-purpose certificates that typically require a revocation mechanism, or short-lived certificates that may optionally omit revocation because they expire quickly.")
		acmPca_createCertificateAuthorityCmd.MarkFlagRequired("certificate-authority-configuration")
		acmPca_createCertificateAuthorityCmd.MarkFlagRequired("certificate-authority-type")
	})
	acmPcaCmd.AddCommand(acmPca_createCertificateAuthorityCmd)
}
