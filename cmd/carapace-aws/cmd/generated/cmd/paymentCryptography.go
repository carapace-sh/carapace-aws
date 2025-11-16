package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var paymentCryptographyCmd = &cobra.Command{
	Use:   "payment-cryptography",
	Short: "Amazon Web Services Payment Cryptography Control Plane APIs manage encryption keys for use during payment-related cryptographic operations.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(paymentCryptographyCmd).Standalone()

	rootCmd.AddCommand(paymentCryptographyCmd)
}
