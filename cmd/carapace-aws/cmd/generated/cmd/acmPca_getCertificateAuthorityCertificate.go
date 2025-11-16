package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var acmPca_getCertificateAuthorityCertificateCmd = &cobra.Command{
	Use:   "get-certificate-authority-certificate",
	Short: "Retrieves the certificate and certificate chain for your private certificate authority (CA) or one that has been shared with you.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(acmPca_getCertificateAuthorityCertificateCmd).Standalone()

	acmPca_getCertificateAuthorityCertificateCmd.Flags().String("certificate-authority-arn", "", "The Amazon Resource Name (ARN) of your private CA.")
	acmPca_getCertificateAuthorityCertificateCmd.MarkFlagRequired("certificate-authority-arn")
	acmPcaCmd.AddCommand(acmPca_getCertificateAuthorityCertificateCmd)
}
