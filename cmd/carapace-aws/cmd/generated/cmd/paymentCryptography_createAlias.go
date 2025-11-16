package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var paymentCryptography_createAliasCmd = &cobra.Command{
	Use:   "create-alias",
	Short: "Creates an *alias*, or a friendly name, for an Amazon Web Services Payment Cryptography key.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(paymentCryptography_createAliasCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(paymentCryptography_createAliasCmd).Standalone()

		paymentCryptography_createAliasCmd.Flags().String("alias-name", "", "A friendly name that you can use to refer to a key.")
		paymentCryptography_createAliasCmd.Flags().String("key-arn", "", "The `KeyARN` of the key to associate with the alias.")
		paymentCryptography_createAliasCmd.MarkFlagRequired("alias-name")
	})
	paymentCryptographyCmd.AddCommand(paymentCryptography_createAliasCmd)
}
