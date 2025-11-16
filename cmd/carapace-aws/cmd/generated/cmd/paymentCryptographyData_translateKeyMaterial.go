package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var paymentCryptographyData_translateKeyMaterialCmd = &cobra.Command{
	Use:   "translate-key-material",
	Short: "Translates an encryption key between different wrapping keys without importing the key into Amazon Web Services Payment Cryptography.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(paymentCryptographyData_translateKeyMaterialCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(paymentCryptographyData_translateKeyMaterialCmd).Standalone()

		paymentCryptographyData_translateKeyMaterialCmd.Flags().String("incoming-key-material", "", "Parameter information of the TR31WrappedKeyBlock containing the transaction key.")
		paymentCryptographyData_translateKeyMaterialCmd.Flags().String("key-check-value-algorithm", "", "The key check value (KCV) algorithm used for calculating the KCV.")
		paymentCryptographyData_translateKeyMaterialCmd.Flags().String("outgoing-key-material", "", "Parameter information of the wrapping key used to wrap the transaction key in the outgoing TR31WrappedKeyBlock.")
		paymentCryptographyData_translateKeyMaterialCmd.MarkFlagRequired("incoming-key-material")
		paymentCryptographyData_translateKeyMaterialCmd.MarkFlagRequired("outgoing-key-material")
	})
	paymentCryptographyDataCmd.AddCommand(paymentCryptographyData_translateKeyMaterialCmd)
}
