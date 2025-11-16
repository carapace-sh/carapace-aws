package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var acmPca_restoreCertificateAuthorityCmd = &cobra.Command{
	Use:   "restore-certificate-authority",
	Short: "Restores a certificate authority (CA) that is in the `DELETED` state.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(acmPca_restoreCertificateAuthorityCmd).Standalone()

	acmPca_restoreCertificateAuthorityCmd.Flags().String("certificate-authority-arn", "", "The Amazon Resource Name (ARN) that was returned when you called the [CreateCertificateAuthority](https://docs.aws.amazon.com/privateca/latest/APIReference/API_CreateCertificateAuthority.html) action.")
	acmPca_restoreCertificateAuthorityCmd.MarkFlagRequired("certificate-authority-arn")
	acmPcaCmd.AddCommand(acmPca_restoreCertificateAuthorityCmd)
}
