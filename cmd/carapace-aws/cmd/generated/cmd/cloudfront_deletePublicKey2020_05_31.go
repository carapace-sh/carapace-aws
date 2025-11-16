package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudfront_deletePublicKey2020_05_31Cmd = &cobra.Command{
	Use:   "delete-public-key2020_05_31",
	Short: "Remove a public key you previously added to CloudFront.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudfront_deletePublicKey2020_05_31Cmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloudfront_deletePublicKey2020_05_31Cmd).Standalone()

		cloudfront_deletePublicKey2020_05_31Cmd.Flags().String("id", "", "The ID of the public key you want to remove from CloudFront.")
		cloudfront_deletePublicKey2020_05_31Cmd.Flags().String("if-match", "", "The value of the `ETag` header that you received when retrieving the public key identity to delete.")
		cloudfront_deletePublicKey2020_05_31Cmd.MarkFlagRequired("id")
	})
	cloudfrontCmd.AddCommand(cloudfront_deletePublicKey2020_05_31Cmd)
}
