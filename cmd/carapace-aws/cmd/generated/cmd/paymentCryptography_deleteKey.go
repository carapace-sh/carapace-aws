package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var paymentCryptography_deleteKeyCmd = &cobra.Command{
	Use:   "delete-key",
	Short: "Deletes the key material and metadata associated with Amazon Web Services Payment Cryptography key.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(paymentCryptography_deleteKeyCmd).Standalone()

	paymentCryptography_deleteKeyCmd.Flags().String("delete-key-in-days", "", "The waiting period for key deletion.")
	paymentCryptography_deleteKeyCmd.Flags().String("key-identifier", "", "The `KeyARN` of the key that is scheduled for deletion.")
	paymentCryptography_deleteKeyCmd.MarkFlagRequired("key-identifier")
	paymentCryptographyCmd.AddCommand(paymentCryptography_deleteKeyCmd)
}
