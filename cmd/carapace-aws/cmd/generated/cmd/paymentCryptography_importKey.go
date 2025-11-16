package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var paymentCryptography_importKeyCmd = &cobra.Command{
	Use:   "import-key",
	Short: "Imports symmetric keys and public key certificates in PEM format (base64 encoded) into Amazon Web Services Payment Cryptography.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(paymentCryptography_importKeyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(paymentCryptography_importKeyCmd).Standalone()

		paymentCryptography_importKeyCmd.Flags().Bool("enabled", false, "Specifies whether import key is enabled.")
		paymentCryptography_importKeyCmd.Flags().String("key-check-value-algorithm", "", "The algorithm that Amazon Web Services Payment Cryptography uses to calculate the key check value (KCV).")
		paymentCryptography_importKeyCmd.Flags().String("key-material", "", "The key or public key certificate type to use during key material import, for example TR-34 or RootCertificatePublicKey.")
		paymentCryptography_importKeyCmd.Flags().Bool("no-enabled", false, "Specifies whether import key is enabled.")
		paymentCryptography_importKeyCmd.Flags().String("replication-regions", "", "")
		paymentCryptography_importKeyCmd.Flags().String("tags", "", "Assigns one or more tags to the Amazon Web Services Payment Cryptography key.")
		paymentCryptography_importKeyCmd.MarkFlagRequired("key-material")
		paymentCryptography_importKeyCmd.Flag("no-enabled").Hidden = true
	})
	paymentCryptographyCmd.AddCommand(paymentCryptography_importKeyCmd)
}
