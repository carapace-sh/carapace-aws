package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var paymentCryptography_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Deletes a tag from an Amazon Web Services Payment Cryptography key.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(paymentCryptography_untagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(paymentCryptography_untagResourceCmd).Standalone()

		paymentCryptography_untagResourceCmd.Flags().String("resource-arn", "", "The `KeyARN` of the key whose tags are being removed.")
		paymentCryptography_untagResourceCmd.Flags().String("tag-keys", "", "One or more tag keys.")
		paymentCryptography_untagResourceCmd.MarkFlagRequired("resource-arn")
		paymentCryptography_untagResourceCmd.MarkFlagRequired("tag-keys")
	})
	paymentCryptographyCmd.AddCommand(paymentCryptography_untagResourceCmd)
}
