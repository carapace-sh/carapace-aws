package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var paymentCryptography_restoreKeyCmd = &cobra.Command{
	Use:   "restore-key",
	Short: "Cancels a scheduled key deletion during the waiting period.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(paymentCryptography_restoreKeyCmd).Standalone()

	paymentCryptography_restoreKeyCmd.Flags().String("key-identifier", "", "The `KeyARN` of the key to be restored within Amazon Web Services Payment Cryptography.")
	paymentCryptography_restoreKeyCmd.MarkFlagRequired("key-identifier")
	paymentCryptographyCmd.AddCommand(paymentCryptography_restoreKeyCmd)
}
