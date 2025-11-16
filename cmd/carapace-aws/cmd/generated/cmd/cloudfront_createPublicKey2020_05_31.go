package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudfront_createPublicKey2020_05_31Cmd = &cobra.Command{
	Use:   "create-public-key2020_05_31",
	Short: "Uploads a public key to CloudFront that you can use with [signed URLs and signed cookies](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/PrivateContent.html), or with [field-level encryption](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/field-level-encryption.html).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudfront_createPublicKey2020_05_31Cmd).Standalone()

	cloudfront_createPublicKey2020_05_31Cmd.Flags().String("public-key-config", "", "A CloudFront public key configuration.")
	cloudfront_createPublicKey2020_05_31Cmd.MarkFlagRequired("public-key-config")
	cloudfrontCmd.AddCommand(cloudfront_createPublicKey2020_05_31Cmd)
}
