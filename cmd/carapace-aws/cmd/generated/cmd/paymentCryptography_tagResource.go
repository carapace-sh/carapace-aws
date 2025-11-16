package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var paymentCryptography_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Adds or edits tags on an Amazon Web Services Payment Cryptography key.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(paymentCryptography_tagResourceCmd).Standalone()

	paymentCryptography_tagResourceCmd.Flags().String("resource-arn", "", "The `KeyARN` of the key whose tags are being updated.")
	paymentCryptography_tagResourceCmd.Flags().String("tags", "", "One or more tags.")
	paymentCryptography_tagResourceCmd.MarkFlagRequired("resource-arn")
	paymentCryptography_tagResourceCmd.MarkFlagRequired("tags")
	paymentCryptographyCmd.AddCommand(paymentCryptography_tagResourceCmd)
}
