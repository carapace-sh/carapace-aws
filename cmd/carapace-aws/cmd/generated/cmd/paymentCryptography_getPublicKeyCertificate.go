package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var paymentCryptography_getPublicKeyCertificateCmd = &cobra.Command{
	Use:   "get-public-key-certificate",
	Short: "Gets the public key certificate of the asymmetric key pair that exists within Amazon Web Services Payment Cryptography.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(paymentCryptography_getPublicKeyCertificateCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(paymentCryptography_getPublicKeyCertificateCmd).Standalone()

		paymentCryptography_getPublicKeyCertificateCmd.Flags().String("key-identifier", "", "The `KeyARN` of the asymmetric key pair.")
		paymentCryptography_getPublicKeyCertificateCmd.MarkFlagRequired("key-identifier")
	})
	paymentCryptographyCmd.AddCommand(paymentCryptography_getPublicKeyCertificateCmd)
}
