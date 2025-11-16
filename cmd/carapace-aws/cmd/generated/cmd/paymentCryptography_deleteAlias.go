package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var paymentCryptography_deleteAliasCmd = &cobra.Command{
	Use:   "delete-alias",
	Short: "Deletes the alias, but doesn't affect the underlying key.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(paymentCryptography_deleteAliasCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(paymentCryptography_deleteAliasCmd).Standalone()

		paymentCryptography_deleteAliasCmd.Flags().String("alias-name", "", "A friendly name that you can use to refer Amazon Web Services Payment Cryptography key.")
		paymentCryptography_deleteAliasCmd.MarkFlagRequired("alias-name")
	})
	paymentCryptographyCmd.AddCommand(paymentCryptography_deleteAliasCmd)
}
