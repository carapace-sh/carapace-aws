package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var paymentCryptography_exportKeyCmd = &cobra.Command{
	Use:   "export-key",
	Short: "Exports a key from Amazon Web Services Payment Cryptography.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(paymentCryptography_exportKeyCmd).Standalone()

	paymentCryptography_exportKeyCmd.Flags().String("export-attributes", "", "The attributes for IPEK generation during export.")
	paymentCryptography_exportKeyCmd.Flags().String("export-key-identifier", "", "The `KeyARN` of the key under export from Amazon Web Services Payment Cryptography.")
	paymentCryptography_exportKeyCmd.Flags().String("key-material", "", "The key block format type, for example, TR-34 or TR-31, to use during key material export.")
	paymentCryptography_exportKeyCmd.MarkFlagRequired("export-key-identifier")
	paymentCryptography_exportKeyCmd.MarkFlagRequired("key-material")
	paymentCryptographyCmd.AddCommand(paymentCryptography_exportKeyCmd)
}
