package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var paymentCryptography_getKeyCmd = &cobra.Command{
	Use:   "get-key",
	Short: "Gets the key metadata for an Amazon Web Services Payment Cryptography key, including the immutable and mutable attributes specified when the key was created.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(paymentCryptography_getKeyCmd).Standalone()

	paymentCryptography_getKeyCmd.Flags().String("key-identifier", "", "The `KeyARN` of the Amazon Web Services Payment Cryptography key.")
	paymentCryptography_getKeyCmd.MarkFlagRequired("key-identifier")
	paymentCryptographyCmd.AddCommand(paymentCryptography_getKeyCmd)
}
