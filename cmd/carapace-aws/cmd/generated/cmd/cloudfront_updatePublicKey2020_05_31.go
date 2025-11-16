package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudfront_updatePublicKey2020_05_31Cmd = &cobra.Command{
	Use:   "update-public-key2020_05_31",
	Short: "Update public key information.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudfront_updatePublicKey2020_05_31Cmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloudfront_updatePublicKey2020_05_31Cmd).Standalone()

		cloudfront_updatePublicKey2020_05_31Cmd.Flags().String("id", "", "The identifier of the public key that you are updating.")
		cloudfront_updatePublicKey2020_05_31Cmd.Flags().String("if-match", "", "The value of the `ETag` header that you received when retrieving the public key to update.")
		cloudfront_updatePublicKey2020_05_31Cmd.Flags().String("public-key-config", "", "A public key configuration.")
		cloudfront_updatePublicKey2020_05_31Cmd.MarkFlagRequired("id")
		cloudfront_updatePublicKey2020_05_31Cmd.MarkFlagRequired("public-key-config")
	})
	cloudfrontCmd.AddCommand(cloudfront_updatePublicKey2020_05_31Cmd)
}
