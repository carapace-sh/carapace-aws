package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudfront_publishFunction2020_05_31Cmd = &cobra.Command{
	Use:   "publish-function2020_05_31",
	Short: "Publishes a CloudFront function by copying the function code from the `DEVELOPMENT` stage to `LIVE`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudfront_publishFunction2020_05_31Cmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloudfront_publishFunction2020_05_31Cmd).Standalone()

		cloudfront_publishFunction2020_05_31Cmd.Flags().String("if-match", "", "The current version (`ETag` value) of the function that you are publishing, which you can get using `DescribeFunction`.")
		cloudfront_publishFunction2020_05_31Cmd.Flags().String("name", "", "The name of the function that you are publishing.")
		cloudfront_publishFunction2020_05_31Cmd.MarkFlagRequired("if-match")
		cloudfront_publishFunction2020_05_31Cmd.MarkFlagRequired("name")
	})
	cloudfrontCmd.AddCommand(cloudfront_publishFunction2020_05_31Cmd)
}
