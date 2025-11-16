package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudfront_getPublicKey2020_05_31Cmd = &cobra.Command{
	Use:   "get-public-key2020_05_31",
	Short: "Gets a public key.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudfront_getPublicKey2020_05_31Cmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloudfront_getPublicKey2020_05_31Cmd).Standalone()

		cloudfront_getPublicKey2020_05_31Cmd.Flags().String("id", "", "The identifier of the public key you are getting.")
		cloudfront_getPublicKey2020_05_31Cmd.MarkFlagRequired("id")
	})
	cloudfrontCmd.AddCommand(cloudfront_getPublicKey2020_05_31Cmd)
}
