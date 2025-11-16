package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var paymentCryptography_getParametersForExportCmd = &cobra.Command{
	Use:   "get-parameters-for-export",
	Short: "Gets the export token and the signing key certificate to initiate a TR-34 key export from Amazon Web Services Payment Cryptography.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(paymentCryptography_getParametersForExportCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(paymentCryptography_getParametersForExportCmd).Standalone()

		paymentCryptography_getParametersForExportCmd.Flags().String("key-material-type", "", "The key block format type (for example, TR-34 or TR-31) to use during key material export.")
		paymentCryptography_getParametersForExportCmd.Flags().String("signing-key-algorithm", "", "The signing key algorithm to generate a signing key certificate.")
		paymentCryptography_getParametersForExportCmd.MarkFlagRequired("key-material-type")
		paymentCryptography_getParametersForExportCmd.MarkFlagRequired("signing-key-algorithm")
	})
	paymentCryptographyCmd.AddCommand(paymentCryptography_getParametersForExportCmd)
}
