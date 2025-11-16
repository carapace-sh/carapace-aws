package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var paymentCryptography_getCertificateSigningRequestCmd = &cobra.Command{
	Use:   "get-certificate-signing-request",
	Short: "Creates a certificate signing request (CSR) from a key pair.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(paymentCryptography_getCertificateSigningRequestCmd).Standalone()

	paymentCryptography_getCertificateSigningRequestCmd.Flags().String("certificate-subject", "", "The metadata used to create the CSR.")
	paymentCryptography_getCertificateSigningRequestCmd.Flags().String("key-identifier", "", "Asymmetric key used for generating the certificate signing request")
	paymentCryptography_getCertificateSigningRequestCmd.Flags().String("signing-algorithm", "", "The cryptographic algorithm used to sign your CSR.")
	paymentCryptography_getCertificateSigningRequestCmd.MarkFlagRequired("certificate-subject")
	paymentCryptography_getCertificateSigningRequestCmd.MarkFlagRequired("key-identifier")
	paymentCryptography_getCertificateSigningRequestCmd.MarkFlagRequired("signing-algorithm")
	paymentCryptographyCmd.AddCommand(paymentCryptography_getCertificateSigningRequestCmd)
}
