package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var paymentCryptography_startKeyUsageCmd = &cobra.Command{
	Use:   "start-key-usage",
	Short: "Enables an Amazon Web Services Payment Cryptography key, which makes it active for cryptographic operations within Amazon Web Services Payment Cryptography\n\n**Cross-account use:** This operation can't be used across different Amazon Web Services accounts.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(paymentCryptography_startKeyUsageCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(paymentCryptography_startKeyUsageCmd).Standalone()

		paymentCryptography_startKeyUsageCmd.Flags().String("key-identifier", "", "The `KeyArn` of the key.")
		paymentCryptography_startKeyUsageCmd.MarkFlagRequired("key-identifier")
	})
	paymentCryptographyCmd.AddCommand(paymentCryptography_startKeyUsageCmd)
}
