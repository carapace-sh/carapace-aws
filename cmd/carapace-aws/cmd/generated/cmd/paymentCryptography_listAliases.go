package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var paymentCryptography_listAliasesCmd = &cobra.Command{
	Use:   "list-aliases",
	Short: "Lists the aliases for all keys in the caller's Amazon Web Services account and Amazon Web Services Region.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(paymentCryptography_listAliasesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(paymentCryptography_listAliasesCmd).Standalone()

		paymentCryptography_listAliasesCmd.Flags().String("key-arn", "", "The `keyARN` for which you want to list all aliases.")
		paymentCryptography_listAliasesCmd.Flags().String("max-results", "", "Use this parameter to specify the maximum number of items to return.")
		paymentCryptography_listAliasesCmd.Flags().String("next-token", "", "Use this parameter in a subsequent request after you receive a response with truncated results.")
	})
	paymentCryptographyCmd.AddCommand(paymentCryptography_listAliasesCmd)
}
