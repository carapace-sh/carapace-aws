package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudfront_getPublicKeyConfig2020_05_31Cmd = &cobra.Command{
	Use:   "get-public-key-config2020_05_31",
	Short: "Gets a public key configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudfront_getPublicKeyConfig2020_05_31Cmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloudfront_getPublicKeyConfig2020_05_31Cmd).Standalone()

		cloudfront_getPublicKeyConfig2020_05_31Cmd.Flags().String("id", "", "The identifier of the public key whose configuration you are getting.")
		cloudfront_getPublicKeyConfig2020_05_31Cmd.MarkFlagRequired("id")
	})
	cloudfrontCmd.AddCommand(cloudfront_getPublicKeyConfig2020_05_31Cmd)
}
