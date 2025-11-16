package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var acmPca_describeCertificateAuthorityCmd = &cobra.Command{
	Use:   "describe-certificate-authority",
	Short: "Lists information about your private certificate authority (CA) or one that has been shared with you.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(acmPca_describeCertificateAuthorityCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(acmPca_describeCertificateAuthorityCmd).Standalone()

		acmPca_describeCertificateAuthorityCmd.Flags().String("certificate-authority-arn", "", "The Amazon Resource Name (ARN) that was returned when you called [CreateCertificateAuthority](https://docs.aws.amazon.com/privateca/latest/APIReference/API_CreateCertificateAuthority.html).")
		acmPca_describeCertificateAuthorityCmd.MarkFlagRequired("certificate-authority-arn")
	})
	acmPcaCmd.AddCommand(acmPca_describeCertificateAuthorityCmd)
}
