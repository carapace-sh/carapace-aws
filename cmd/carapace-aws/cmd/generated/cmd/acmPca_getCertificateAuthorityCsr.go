package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var acmPca_getCertificateAuthorityCsrCmd = &cobra.Command{
	Use:   "get-certificate-authority-csr",
	Short: "Retrieves the certificate signing request (CSR) for your private certificate authority (CA).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(acmPca_getCertificateAuthorityCsrCmd).Standalone()

	acmPca_getCertificateAuthorityCsrCmd.Flags().String("certificate-authority-arn", "", "The Amazon Resource Name (ARN) that was returned when you called the [CreateCertificateAuthority](https://docs.aws.amazon.com/privateca/latest/APIReference/API_CreateCertificateAuthority.html) action.")
	acmPca_getCertificateAuthorityCsrCmd.MarkFlagRequired("certificate-authority-arn")
	acmPcaCmd.AddCommand(acmPca_getCertificateAuthorityCsrCmd)
}
