package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var paymentCryptography_stopKeyUsageCmd = &cobra.Command{
	Use:   "stop-key-usage",
	Short: "Disables an Amazon Web Services Payment Cryptography key, which makes it inactive within Amazon Web Services Payment Cryptography.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(paymentCryptography_stopKeyUsageCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(paymentCryptography_stopKeyUsageCmd).Standalone()

		paymentCryptography_stopKeyUsageCmd.Flags().String("key-identifier", "", "The `KeyArn` of the key.")
		paymentCryptography_stopKeyUsageCmd.MarkFlagRequired("key-identifier")
	})
	paymentCryptographyCmd.AddCommand(paymentCryptography_stopKeyUsageCmd)
}
