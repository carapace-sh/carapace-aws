package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudfront_getFunction2020_05_31Cmd = &cobra.Command{
	Use:   "get-function2020_05_31",
	Short: "Gets the code of a CloudFront function.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudfront_getFunction2020_05_31Cmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloudfront_getFunction2020_05_31Cmd).Standalone()

		cloudfront_getFunction2020_05_31Cmd.Flags().String("name", "", "The name of the function whose code you are getting.")
		cloudfront_getFunction2020_05_31Cmd.Flags().String("stage", "", "The function's stage, either `DEVELOPMENT` or `LIVE`.")
		cloudfront_getFunction2020_05_31Cmd.MarkFlagRequired("name")
	})
	cloudfrontCmd.AddCommand(cloudfront_getFunction2020_05_31Cmd)
}
