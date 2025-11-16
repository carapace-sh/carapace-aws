package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var paymentCryptographyDataCmd = &cobra.Command{
	Use:   "payment-cryptography-data",
	Short: "You use the Amazon Web Services Payment Cryptography Data Plane to manage how encryption keys are used for payment-related transaction processing and associated cryptographic operations.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(paymentCryptographyDataCmd).Standalone()

	rootCmd.AddCommand(paymentCryptographyDataCmd)
}
