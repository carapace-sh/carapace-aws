package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var acmPca_importCertificateAuthorityCertificateCmd = &cobra.Command{
	Use:   "import-certificate-authority-certificate",
	Short: "Imports a signed private CA certificate into Amazon Web Services Private CA.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(acmPca_importCertificateAuthorityCertificateCmd).Standalone()

	acmPca_importCertificateAuthorityCertificateCmd.Flags().String("certificate", "", "The PEM-encoded certificate for a private CA.")
	acmPca_importCertificateAuthorityCertificateCmd.Flags().String("certificate-authority-arn", "", "The Amazon Resource Name (ARN) that was returned when you called [CreateCertificateAuthority](https://docs.aws.amazon.com/privateca/latest/APIReference/API_CreateCertificateAuthority.html).")
	acmPca_importCertificateAuthorityCertificateCmd.Flags().String("certificate-chain", "", "A PEM-encoded file that contains all of your certificates, other than the certificate you're importing, chaining up to your root CA.")
	acmPca_importCertificateAuthorityCertificateCmd.MarkFlagRequired("certificate")
	acmPca_importCertificateAuthorityCertificateCmd.MarkFlagRequired("certificate-authority-arn")
	acmPcaCmd.AddCommand(acmPca_importCertificateAuthorityCertificateCmd)
}
