package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var paymentCryptography_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "Lists the tags for an Amazon Web Services resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(paymentCryptography_listTagsForResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(paymentCryptography_listTagsForResourceCmd).Standalone()

		paymentCryptography_listTagsForResourceCmd.Flags().String("max-results", "", "Use this parameter to specify the maximum number of items to return.")
		paymentCryptography_listTagsForResourceCmd.Flags().String("next-token", "", "Use this parameter in a subsequent request after you receive a response with truncated results.")
		paymentCryptography_listTagsForResourceCmd.Flags().String("resource-arn", "", "The `KeyARN` of the key whose tags you are getting.")
		paymentCryptography_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	})
	paymentCryptographyCmd.AddCommand(paymentCryptography_listTagsForResourceCmd)
}
