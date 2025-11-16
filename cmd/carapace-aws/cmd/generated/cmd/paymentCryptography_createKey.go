package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var paymentCryptography_createKeyCmd = &cobra.Command{
	Use:   "create-key",
	Short: "Creates an Amazon Web Services Payment Cryptography key, a logical representation of a cryptographic key, that is unique in your account and Amazon Web Services Region.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(paymentCryptography_createKeyCmd).Standalone()

	paymentCryptography_createKeyCmd.Flags().String("derive-key-usage", "", "The intended cryptographic usage of keys derived from the ECC key pair to be created.")
	paymentCryptography_createKeyCmd.Flags().Bool("enabled", false, "Specifies whether to enable the key.")
	paymentCryptography_createKeyCmd.Flags().Bool("exportable", false, "Specifies whether the key is exportable from the service.")
	paymentCryptography_createKeyCmd.Flags().String("key-attributes", "", "The role of the key, the algorithm it supports, and the cryptographic operations allowed with the key.")
	paymentCryptography_createKeyCmd.Flags().String("key-check-value-algorithm", "", "The algorithm that Amazon Web Services Payment Cryptography uses to calculate the key check value (KCV).")
	paymentCryptography_createKeyCmd.Flags().Bool("no-enabled", false, "Specifies whether to enable the key.")
	paymentCryptography_createKeyCmd.Flags().Bool("no-exportable", false, "Specifies whether the key is exportable from the service.")
	paymentCryptography_createKeyCmd.Flags().String("replication-regions", "", "")
	paymentCryptography_createKeyCmd.Flags().String("tags", "", "Assigns one or more tags to the Amazon Web Services Payment Cryptography key.")
	paymentCryptography_createKeyCmd.MarkFlagRequired("exportable")
	paymentCryptography_createKeyCmd.MarkFlagRequired("key-attributes")
	paymentCryptography_createKeyCmd.Flag("no-enabled").Hidden = true
	paymentCryptography_createKeyCmd.Flag("no-exportable").Hidden = true
	paymentCryptography_createKeyCmd.MarkFlagRequired("no-exportable")
	paymentCryptographyCmd.AddCommand(paymentCryptography_createKeyCmd)
}
