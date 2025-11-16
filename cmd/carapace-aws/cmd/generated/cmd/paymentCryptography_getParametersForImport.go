package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var paymentCryptography_getParametersForImportCmd = &cobra.Command{
	Use:   "get-parameters-for-import",
	Short: "Gets the import token and the wrapping key certificate in PEM format (base64 encoded) to initiate a TR-34 WrappedKeyBlock or a RSA WrappedKeyCryptogram import into Amazon Web Services Payment Cryptography.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(paymentCryptography_getParametersForImportCmd).Standalone()

	paymentCryptography_getParametersForImportCmd.Flags().String("key-material-type", "", "The method to use for key material import.")
	paymentCryptography_getParametersForImportCmd.Flags().String("wrapping-key-algorithm", "", "The wrapping key algorithm to generate a wrapping key certificate.")
	paymentCryptography_getParametersForImportCmd.MarkFlagRequired("key-material-type")
	paymentCryptography_getParametersForImportCmd.MarkFlagRequired("wrapping-key-algorithm")
	paymentCryptographyCmd.AddCommand(paymentCryptography_getParametersForImportCmd)
}
