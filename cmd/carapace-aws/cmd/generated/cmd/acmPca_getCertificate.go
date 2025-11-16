package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var acmPca_getCertificateCmd = &cobra.Command{
	Use:   "get-certificate",
	Short: "Retrieves a certificate from your private CA or one that has been shared with you.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(acmPca_getCertificateCmd).Standalone()

	acmPca_getCertificateCmd.Flags().String("certificate-arn", "", "The ARN of the issued certificate.")
	acmPca_getCertificateCmd.Flags().String("certificate-authority-arn", "", "The Amazon Resource Name (ARN) that was returned when you called [CreateCertificateAuthority](https://docs.aws.amazon.com/privateca/latest/APIReference/API_CreateCertificateAuthority.html).")
	acmPca_getCertificateCmd.MarkFlagRequired("certificate-arn")
	acmPca_getCertificateCmd.MarkFlagRequired("certificate-authority-arn")
	acmPcaCmd.AddCommand(acmPca_getCertificateCmd)
}
