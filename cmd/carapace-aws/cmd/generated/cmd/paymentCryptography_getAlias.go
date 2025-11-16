package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var paymentCryptography_getAliasCmd = &cobra.Command{
	Use:   "get-alias",
	Short: "Gets the Amazon Web Services Payment Cryptography key associated with the alias.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(paymentCryptography_getAliasCmd).Standalone()

	paymentCryptography_getAliasCmd.Flags().String("alias-name", "", "The alias of the Amazon Web Services Payment Cryptography key.")
	paymentCryptography_getAliasCmd.MarkFlagRequired("alias-name")
	paymentCryptographyCmd.AddCommand(paymentCryptography_getAliasCmd)
}
