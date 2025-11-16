package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var paymentCryptography_updateAliasCmd = &cobra.Command{
	Use:   "update-alias",
	Short: "Associates an existing Amazon Web Services Payment Cryptography alias with a different key.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(paymentCryptography_updateAliasCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(paymentCryptography_updateAliasCmd).Standalone()

		paymentCryptography_updateAliasCmd.Flags().String("alias-name", "", "The alias whose associated key is changing.")
		paymentCryptography_updateAliasCmd.Flags().String("key-arn", "", "The `KeyARN` for the key that you are updating or removing from the alias.")
		paymentCryptography_updateAliasCmd.MarkFlagRequired("alias-name")
	})
	paymentCryptographyCmd.AddCommand(paymentCryptography_updateAliasCmd)
}
