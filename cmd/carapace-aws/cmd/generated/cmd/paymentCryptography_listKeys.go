package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var paymentCryptography_listKeysCmd = &cobra.Command{
	Use:   "list-keys",
	Short: "Lists the keys in the caller's Amazon Web Services account and Amazon Web Services Region.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(paymentCryptography_listKeysCmd).Standalone()

	paymentCryptography_listKeysCmd.Flags().String("key-state", "", "The key state of the keys you want to list.")
	paymentCryptography_listKeysCmd.Flags().String("max-results", "", "Use this parameter to specify the maximum number of items to return.")
	paymentCryptography_listKeysCmd.Flags().String("next-token", "", "Use this parameter in a subsequent request after you receive a response with truncated results.")
	paymentCryptographyCmd.AddCommand(paymentCryptography_listKeysCmd)
}
