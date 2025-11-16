package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var acmPca_revokeCertificateCmd = &cobra.Command{
	Use:   "revoke-certificate",
	Short: "Revokes a certificate that was issued inside Amazon Web Services Private CA.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(acmPca_revokeCertificateCmd).Standalone()

	acmPca_revokeCertificateCmd.Flags().String("certificate-authority-arn", "", "Amazon Resource Name (ARN) of the private CA that issued the certificate to be revoked.")
	acmPca_revokeCertificateCmd.Flags().String("certificate-serial", "", "Serial number of the certificate to be revoked.")
	acmPca_revokeCertificateCmd.Flags().String("revocation-reason", "", "Specifies why you revoked the certificate.")
	acmPca_revokeCertificateCmd.MarkFlagRequired("certificate-authority-arn")
	acmPca_revokeCertificateCmd.MarkFlagRequired("certificate-serial")
	acmPca_revokeCertificateCmd.MarkFlagRequired("revocation-reason")
	acmPcaCmd.AddCommand(acmPca_revokeCertificateCmd)
}
