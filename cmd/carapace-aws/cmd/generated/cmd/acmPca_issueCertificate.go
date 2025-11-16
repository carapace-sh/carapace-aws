package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var acmPca_issueCertificateCmd = &cobra.Command{
	Use:   "issue-certificate",
	Short: "Uses your private certificate authority (CA), or one that has been shared with you, to issue a client certificate.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(acmPca_issueCertificateCmd).Standalone()

	acmPca_issueCertificateCmd.Flags().String("api-passthrough", "", "Specifies X.509 certificate information to be included in the issued certificate.")
	acmPca_issueCertificateCmd.Flags().String("certificate-authority-arn", "", "The Amazon Resource Name (ARN) that was returned when you called [CreateCertificateAuthority](https://docs.aws.amazon.com/privateca/latest/APIReference/API_CreateCertificateAuthority.html).")
	acmPca_issueCertificateCmd.Flags().String("csr", "", "The certificate signing request (CSR) for the certificate you want to issue.")
	acmPca_issueCertificateCmd.Flags().String("idempotency-token", "", "Alphanumeric string that can be used to distinguish between calls to the **IssueCertificate** action.")
	acmPca_issueCertificateCmd.Flags().String("signing-algorithm", "", "The name of the algorithm that will be used to sign the certificate to be issued.")
	acmPca_issueCertificateCmd.Flags().String("template-arn", "", "Specifies a custom configuration template to use when issuing a certificate.")
	acmPca_issueCertificateCmd.Flags().String("validity", "", "Information describing the end of the validity period of the certificate.")
	acmPca_issueCertificateCmd.Flags().String("validity-not-before", "", "Information describing the start of the validity period of the certificate.")
	acmPca_issueCertificateCmd.MarkFlagRequired("certificate-authority-arn")
	acmPca_issueCertificateCmd.MarkFlagRequired("csr")
	acmPca_issueCertificateCmd.MarkFlagRequired("signing-algorithm")
	acmPca_issueCertificateCmd.MarkFlagRequired("validity")
	acmPcaCmd.AddCommand(acmPca_issueCertificateCmd)
}
