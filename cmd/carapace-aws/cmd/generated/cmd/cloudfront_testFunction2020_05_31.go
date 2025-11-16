package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudfront_testFunction2020_05_31Cmd = &cobra.Command{
	Use:   "test-function2020_05_31",
	Short: "Tests a CloudFront function.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudfront_testFunction2020_05_31Cmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloudfront_testFunction2020_05_31Cmd).Standalone()

		cloudfront_testFunction2020_05_31Cmd.Flags().String("event-object", "", "The event object to test the function with.")
		cloudfront_testFunction2020_05_31Cmd.Flags().String("if-match", "", "The current version (`ETag` value) of the function that you are testing, which you can get using `DescribeFunction`.")
		cloudfront_testFunction2020_05_31Cmd.Flags().String("name", "", "The name of the function that you are testing.")
		cloudfront_testFunction2020_05_31Cmd.Flags().String("stage", "", "The stage of the function that you are testing, either `DEVELOPMENT` or `LIVE`.")
		cloudfront_testFunction2020_05_31Cmd.MarkFlagRequired("event-object")
		cloudfront_testFunction2020_05_31Cmd.MarkFlagRequired("if-match")
		cloudfront_testFunction2020_05_31Cmd.MarkFlagRequired("name")
	})
	cloudfrontCmd.AddCommand(cloudfront_testFunction2020_05_31Cmd)
}
